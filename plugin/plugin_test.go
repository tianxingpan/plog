package plugin_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tianxingpan/plog/plugin"
)

const (
	pluginType        = "mock_type"
	pluginName        = "mock_name"
	pluginFailName    = "mock_fail_name"
	pluginTimeoutName = "mock_timeout_name"
	pluginDependName  = "mock_depend_name"
)

type mockPlugin struct{}

func (p *mockPlugin) Type() string {
	return pluginType
}

func (p *mockPlugin) Setup(name string, decoder plugin.Decoder) error {
	return nil
}

func TestGet(t *testing.T) {
	plugin.Register(pluginName, &mockPlugin{})
	// test duplicate registration
	plugin.Register(pluginName, &mockPlugin{})
	p := plugin.Get(pluginType, pluginName)
	assert.NotNil(t, p)

	pNo := plugin.Get("notexist", pluginName)
	assert.Nil(t, pNo)
}
