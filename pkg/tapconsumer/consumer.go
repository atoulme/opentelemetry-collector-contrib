// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tapconsumer // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/tapconsumer"
import (
	"go.opentelemetry.io/collector/component"
)

// Consumer registers taps to expose them.
type Consumer interface {
	// RegisterTap registers tap information on the consumer
	RegisterTap(name component.ID, endpoint string)
}
