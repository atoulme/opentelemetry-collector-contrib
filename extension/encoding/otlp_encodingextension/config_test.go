// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlpencodingextension

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	c := &Config{}
	assert.Error(t, c.Validate(), "ex")
	c.Protocol = OTLPProtocolJSON
	assert.NoError(t, c.Validate())
}
