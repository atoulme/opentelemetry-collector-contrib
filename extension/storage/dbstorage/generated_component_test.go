// Code generated by mdatagen. DO NOT EDIT.

package dbstorageextension

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"

	"go.opentelemetry.io/collector/extension/extensiontest"

	"go.opentelemetry.io/collector/confmap/confmaptest"
)

// assertNoErrorHost implements a component.Host that asserts that there were no errors.
type assertNoErrorHost struct {
	component.Host
	*testing.T
}

var _ component.Host = (*assertNoErrorHost)(nil)

// newAssertNoErrorHost returns a new instance of assertNoErrorHost.
func newAssertNoErrorHost(t *testing.T) component.Host {
	return &assertNoErrorHost{
		componenttest.NewNopHost(),
		t,
	}
}

func (aneh *assertNoErrorHost) ReportFatalError(err error) {
	assert.NoError(aneh, err)
}

func Test_ComponentLifecycle(t *testing.T) {
	factory := NewFactory()

	cm, err := confmaptest.LoadConf("metadata.yaml")
	require.NoError(t, err)
	cfg := factory.CreateDefaultConfig()
	sub, err := cm.Sub("tests::config")
	require.NoError(t, err)
	require.NoError(t, component.UnmarshalConfig(sub, cfg))

	t.Run("shutdown", func(t *testing.T) {
		e, err := factory.CreateExtension(context.Background(), extensiontest.NewNopCreateSettings(), cfg)
		require.NoError(t, err)
		err = e.Shutdown(context.Background())
		require.NoError(t, err)
	})

	t.Run("lifecycle", func(t *testing.T) {

		firstExt, err := factory.CreateExtension(context.Background(), extensiontest.NewNopCreateSettings(), cfg)
		require.NoError(t, err)
		require.NoError(t, firstExt.Start(context.Background(), newAssertNoErrorHost(t)))
		require.NoError(t, firstExt.Shutdown(context.Background()))

		secondExt, err := factory.CreateExtension(context.Background(), extensiontest.NewNopCreateSettings(), cfg)
		require.NoError(t, err)
		require.NoError(t, secondExt.Start(context.Background(), newAssertNoErrorHost(t)))
		require.NoError(t, secondExt.Shutdown(context.Background()))
	})
}
