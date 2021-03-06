// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package signalfxexporter

import (
	"net/url"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.uber.org/zap"
)

func TestLoadConfig(t *testing.T) {
	factories, err := config.ExampleComponents()
	assert.Nil(t, err)

	factory := &Factory{}
	factories.Exporters[configmodels.Type(typeStr)] = factory
	cfg, err := config.LoadConfigFile(t, path.Join(".", "testdata", "config.yaml"), factories)

	require.NoError(t, err)
	require.NotNil(t, cfg)

	e0 := cfg.Exporters["signalfx"]

	// Realm doesn't have a default value so set it directly.
	defaultCfg := factory.CreateDefaultConfig().(*Config)
	defaultCfg.Realm = "ap0"
	assert.Equal(t, defaultCfg, e0)

	expectedName := "signalfx/allsettings"

	e1 := cfg.Exporters[expectedName]
	expectedCfg := Config{
		ExporterSettings: configmodels.ExporterSettings{
			TypeVal: configmodels.Type(typeStr),
			NameVal: expectedName,
		},
		AccessToken: "testToken",
		Realm:       "us1",
		Headers: map[string]string{
			"added-entry": "added value",
			"dot.test":    "test",
		},
		Timeout: 2 * time.Second,
	}
	assert.Equal(t, &expectedCfg, e1)

	te, err := factory.CreateMetricsExporter(zap.NewNop(), e1)
	require.NoError(t, err)
	require.NotNil(t, te)
}

func TestConfig_getOptionsFromConfig(t *testing.T) {
	type fields struct {
		ExporterSettings configmodels.ExporterSettings
		AccessToken      string
		Realm            string
		IngestURL        string
		Timeout          time.Duration
		Headers          map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *exporterOptions
		wantErr bool
	}{
		{
			name: "Test URL overrides",
			fields: fields{
				Realm:     "us0",
				IngestURL: "https://ingest.us1.signalfx.com/",
			},
			want: &exporterOptions{
				ingestURL: &url.URL{
					Scheme: "https",
					Host:   "ingest.us1.signalfx.com",
					Path:   "/v2/datapoint",
				},
				httpTimeout: 5 * time.Second,
			},
			wantErr: false,
		},
		{
			name: "Test URL from Realm",
			fields: fields{
				Realm:   "us0",
				Timeout: 10 * time.Second,
			},
			want: &exporterOptions{
				ingestURL: &url.URL{
					Scheme: "https",
					Host:   "ingest.us0.signalfx.com",
					Path:   "/v2/datapoint",
				},
				httpTimeout: 10 * time.Second,
			},
			wantErr: false,
		},
		{
			name:    "Test empty config",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{
				ExporterSettings: tt.fields.ExporterSettings,
				AccessToken:      tt.fields.AccessToken,
				Realm:            tt.fields.Realm,
				IngestURL:        tt.fields.IngestURL,
				Timeout:          tt.fields.Timeout,
				Headers:          tt.fields.Headers,
			}
			got, err := cfg.getOptionsFromConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("getOptionsFromConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
