// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"

	"gopkg.in/yaml.v3"
)

func main() {
	folder := flag.String("folder", ".", "folder investigated for modules")
	allowlistFilePath := flag.String("allowlist", "cmd/checkapi/allowlist.txt", "path to a file containing an allowlist of paths to ignore")
	flag.Parse()
	if err := run(*folder, *allowlistFilePath); err != nil {
		log.Fatal(err)
	}
}

type function struct {
	Name        string   `json:"name"`
	Receiver    string   `json:"receiver"`
	ReturnTypes []string `json:"return_types,omitempty"`
	ParamTypes  []string `json:"param_types,omitempty"`
}

type apistruct struct {
	Name   string   `json:"name"`
	Fields []string `json:"fields"`
}

type api struct {
	Values    []string     `json:"values,omitempty"`
	Structs   []*apistruct `json:"structs,omitempty"`
	Functions []*function  `json:"functions,omitempty"`
}

type metadata struct {
	Type   string `yaml:"type"`
	Status status `yaml:"status"`
}

type status struct {
	Class string `yaml:"class"`
}

func run(folder string, allowlistFilePath string) error {
	allowlistData, err := os.ReadFile(allowlistFilePath)
	if err != nil {
		return err
	}
	allowlist := strings.Split(string(allowlistData), "\n")
	var errs []error
	err = filepath.Walk(folder, func(path string, info fs.FileInfo, _ error) error {
		if info.Name() == "go.mod" {
			base := filepath.Dir(path)
			var relativeBase string
			relativeBase, err = filepath.Rel(folder, base)
			if err != nil {
				return err
			}
			if strings.HasPrefix(relativeBase, "internal") {
				return nil
			}

			var componentType string
			if _, err = os.Stat(filepath.Join(base, "metadata.yaml")); errors.Is(err, os.ErrNotExist) {
				componentType = "pkg"
			} else {
				m, err := os.ReadFile(filepath.Join(base, "metadata.yaml"))
				if err != nil {
					return err
				}
				var componentInfo metadata
				if err = yaml.Unmarshal(m, &componentInfo); err != nil {
					return err
				}
				componentType = componentInfo.Status.Class
			}

			for _, a := range allowlist {
				if a == relativeBase {
					fmt.Printf("Ignoring %s per allowlist\n", base)
					return nil
				}
			}
			if err = walkFolder(base, componentType); err != nil {
				errs = append(errs, err)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func isTestFunction(fnName string) bool {
	return strings.HasPrefix(fnName, "Test") ||
		strings.HasPrefix(fnName, "Benchmark") ||
		strings.HasPrefix(fnName, "Fuzz")
}

func handleFile(f *ast.File, result *api) {
	for _, d := range f.Decls {
		if str, isStr := d.(*ast.GenDecl); isStr {
			for _, s := range str.Specs {
				if values, ok := s.(*ast.ValueSpec); ok {
					for _, v := range values.Names {
						if v.IsExported() {
							result.Values = append(result.Values, v.Name)
						}
					}
				}
				if t, ok := s.(*ast.TypeSpec); ok {
					var fieldNames []string
					if t.TypeParams != nil {
						fieldNames = make([]string, len(t.TypeParams.List))
						for i, f := range t.TypeParams.List {
							fieldNames[i] = f.Names[0].Name
						}
					}
					result.Structs = append(result.Structs, &apistruct{
						Name:   t.Name.String(),
						Fields: fieldNames,
					})
				}
			}
		}
		if fn, isFn := d.(*ast.FuncDecl); isFn {
			if !fn.Name.IsExported() {
				continue
			}
			exported := false
			receiver := ""
			if fn.Recv.NumFields() == 0 && !isTestFunction(fn.Name.String()) {
				exported = true
			}
			if fn.Recv.NumFields() > 0 {
				for _, t := range fn.Recv.List {
					for _, n := range t.Names {
						exported = exported || n.IsExported()
						if n.IsExported() {
							receiver = n.Name
						}
					}
				}
			}
			if exported {
				var returnTypes []string
				if fn.Type.Results.NumFields() > 0 {
					for _, r := range fn.Type.Results.List {
						returnTypes = append(returnTypes, exprToString(r.Type))
					}
				}
				var params []string
				if fn.Type.Params.NumFields() > 0 {
					for _, r := range fn.Type.Params.List {
						params = append(params, exprToString(r.Type))
					}
				}
				f := &function{
					Name:        fn.Name.Name,
					Receiver:    receiver,
					ParamTypes:  params,
					ReturnTypes: returnTypes,
				}
				result.Functions = append(result.Functions, f)
			}
		}
	}
}

func walkFolder(folder string, componentType string) error {
	result := &api{}
	set := token.NewFileSet()
	packs, err := parser.ParseDir(set, folder, nil, 0)
	if err != nil {
		return err
	}

	for _, pack := range packs {
		for _, f := range pack.Files {
			handleFile(f, result)
		}
	}
	sort.Slice(result.Structs, func(i int, j int) bool {
		return strings.Compare(result.Structs[i].Name, result.Structs[j].Name) > 0
	})
	sort.Strings(result.Values)
	sort.Slice(result.Functions, func(i int, j int) bool {
		return strings.Compare(result.Functions[i].Name, result.Functions[j].Name) < 0
	})
	fnNames := make([]string, len(result.Functions))
	for i, fn := range result.Functions {
		fnNames[i] = fn.Name
	}
	if len(result.Structs) == 0 && len(result.Values) == 0 && len(result.Functions) == 0 {
		return nil
	}

	if len(result.Functions) > 1 && componentType != "pkg" && componentType != "cmd" {
		return fmt.Errorf("%s has more than one function: %q", folder, strings.Join(fnNames, ","))
	}
	if len(result.Functions) == 1 && componentType != "pkg" && componentType != "cmd" {
		if err := checkFactoryFunction(result.Functions[0], folder, componentType); err != nil {
			return err
		}
	}
	for _, s := range result.Structs {
		if err := checkStructDisallowUnkeyedLiteral(s, folder); err != nil {
			return err
		}
	}
	return nil
}

// check the only exported function of the module is NewFactory, matching the signature of the factory expected by the collector builder.
func checkFactoryFunction(newFactoryFn *function, folder string, componentType string) error {
	switch componentType {
	case "provider":
		return checkProviderFactoryFunction(newFactoryFn, folder, componentType)
	default:
		return checkComponentFactoryFunction(newFactoryFn, folder, componentType)
	}
}

func checkProviderFactoryFunction(newFactoryFn *function, folder string, componentType string) error {
	if newFactoryFn.Name != "NewFactory" {
		return fmt.Errorf("%s does not define a NewFactory function as a %s", folder, componentType)
	}
	if newFactoryFn.Receiver != "" {
		return fmt.Errorf("%s associated NewFactory with a receiver type", folder)
	}
	if len(newFactoryFn.ReturnTypes) != 1 {
		return fmt.Errorf("%s NewFactory function returns more than one result", folder)
	}
	returnType := newFactoryFn.ReturnTypes[0]

	if returnType != "confmap.ProviderFactory" {
		return fmt.Errorf("%s NewFactory function does not return a valid type: %s, expected confmap.ProviderFactory", folder, returnType)
	}
	return nil
}

func checkComponentFactoryFunction(newFactoryFn *function, folder string, componentType string) error {
	if newFactoryFn.Name != "NewFactory" {
		return fmt.Errorf("%s does not define a NewFactory function as a %s", folder, componentType)
	}
	if newFactoryFn.Receiver != "" {
		return fmt.Errorf("%s associated NewFactory with a receiver type", folder)
	}
	if len(newFactoryFn.ReturnTypes) != 1 {
		return fmt.Errorf("%s NewFactory function returns more than one result", folder)
	}
	returnType := newFactoryFn.ReturnTypes[0]

	if returnType != fmt.Sprintf("%s.Factory", componentType) {
		return fmt.Errorf("%s NewFactory function does not return a valid type: %s, expected %s.Factory", folder, returnType, componentType)
	}
	return nil
}

func checkStructDisallowUnkeyedLiteral(s *apistruct, folder string) error {
	if !unicode.IsUpper(rune(s.Name[0])) {
		return nil
	}
	for _, f := range s.Fields {
		if !unicode.IsUpper(rune(f[0])) {
			return nil
		}
	}
	return fmt.Errorf("%s struct %q does not prevent unkeyed literal initialization", folder, s.Name)
}

func exprToString(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", exprToString(e.Key), exprToString(e.Value))
	case *ast.ArrayType:
		return fmt.Sprintf("[%s]%s", exprToString(e.Len), exprToString(e.Elt))
	case *ast.StructType:
		var fields []string
		for _, f := range e.Fields.List {
			fields = append(fields, exprToString(f.Type))
		}
		return fmt.Sprintf("{%s}", strings.Join(fields, ","))
	case *ast.InterfaceType:
		var methods []string
		for _, f := range e.Methods.List {
			methods = append(methods, "func "+exprToString(f.Type))
		}
		return fmt.Sprintf("{%s}", strings.Join(methods, ","))
	case *ast.ChanType:
		return fmt.Sprintf("chan(%s)", exprToString(e.Value))
	case *ast.FuncType:
		var results []string
		if e.Results != nil {
			for _, r := range e.Results.List {
				results = append(results, exprToString(r.Type))
			}
		}
		var params []string
		if e.Params != nil {
			for _, r := range e.Params.List {
				params = append(params, exprToString(r.Type))
			}
		}
		return fmt.Sprintf("func(%s) %s", strings.Join(params, ","), strings.Join(results, ","))
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", exprToString(e.X), e.Sel.Name)
	case *ast.Ident:
		return e.Name
	case nil:
		return ""
	case *ast.StarExpr:
		return fmt.Sprintf("*%s", exprToString(e.X))
	case *ast.Ellipsis:
		return fmt.Sprintf("%s...", exprToString(e.Elt))
	case *ast.IndexExpr:
		return fmt.Sprintf("%s[%s]", exprToString(e.X), exprToString(e.Index))
	case *ast.BasicLit:
		return e.Value
	case *ast.IndexListExpr:
		var exprs []string
		for _, e := range e.Indices {
			exprs = append(exprs, exprToString(e))
		}
		return strings.Join(exprs, ",")
	default:
		panic(fmt.Sprintf("Unsupported expr type: %#v", expr))
	}
}
