// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package targetallocatorextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/targetallocatorextension"
import (
	"context"
	
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

type targetAllocatorExtension struct {
	config   *Config
	settings extension.Settings
}

func (t targetAllocatorExtension) Start(ctx context.Context, host component.Host) error {
	//TODO implement me
	panic("implement me")
}

func (t targetAllocatorExtension) Shutdown(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
