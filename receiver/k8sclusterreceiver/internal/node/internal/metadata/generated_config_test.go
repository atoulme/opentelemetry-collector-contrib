// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestMetricsBuilderConfig(t *testing.T) {
	tests := []struct {
		name string
		want MetricsBuilderConfig
	}{
		{
			name: "default",
			want: DefaultMetricsBuilderConfig(),
		},
		{
			name: "all_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					K8sNodeAllocatableCPU:              MetricConfig{Enabled: true},
					K8sNodeAllocatableEphemeralStorage: MetricConfig{Enabled: true},
					K8sNodeAllocatableMemory:           MetricConfig{Enabled: true},
					K8sNodeConditionDiskPressure:       MetricConfig{Enabled: true},
					K8sNodeConditionMemoryPressure:     MetricConfig{Enabled: true},
					K8sNodeConditionNetworkUnavailable: MetricConfig{Enabled: true},
					K8sNodeConditionPidPressure:        MetricConfig{Enabled: true},
					K8sNodeConditionReady:              MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					K8sNodeName:            ResourceAttributeConfig{Enabled: true},
					K8sNodeUID:             ResourceAttributeConfig{Enabled: true},
					OpencensusResourcetype: ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					K8sNodeAllocatableCPU:              MetricConfig{Enabled: false},
					K8sNodeAllocatableEphemeralStorage: MetricConfig{Enabled: false},
					K8sNodeAllocatableMemory:           MetricConfig{Enabled: false},
					K8sNodeConditionDiskPressure:       MetricConfig{Enabled: false},
					K8sNodeConditionMemoryPressure:     MetricConfig{Enabled: false},
					K8sNodeConditionNetworkUnavailable: MetricConfig{Enabled: false},
					K8sNodeConditionPidPressure:        MetricConfig{Enabled: false},
					K8sNodeConditionReady:              MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					K8sNodeName:            ResourceAttributeConfig{Enabled: false},
					K8sNodeUID:             ResourceAttributeConfig{Enabled: false},
					OpencensusResourcetype: ResourceAttributeConfig{Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadMetricsBuilderConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(MetricConfig{}, ResourceAttributeConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadMetricsBuilderConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}
