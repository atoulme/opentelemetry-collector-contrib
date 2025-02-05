// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package targetallocatorextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/targetallocatorextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/targetallocatorextension/internal/metadata"
)

func NewFactory() extension.Factory {
	return extension.NewFactory(
		metadata.Type,
		createDefaultConfig,
		createExtension,
		component.StabilityLevelDevelopment,
	)
}

func createExtension(_ context.Context, settings extension.Settings, config component.Config) (extension.Extension, error) {
	return &targetAllocatorExtension{config: config.(*Config), settings: settings}, nil
}
