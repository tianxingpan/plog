package plog

import (
	"testing"
)

func TestLog(t *testing.T) {
	l := WithFields("uid", "10012")

	l.Trace("helloworld")
	l.Debug("helloworld")
	l.Info("helloworld")
	l.Warn("helloworld")
	l.Error("helloworld")
	l.Tracef("helloworld")
	l.Debugf("helloworld")
	l.Infof("helloworld")
	l.Warnf("helloworld")
	l.Errorf("helloworld")
}

const illConfigInfo = `
log:
  - writer: file # local file log
    level: debug # std output log level
    writer_config: # config of local file output
      filename:  # path of local file rolling log files
      roll_type: time    # rolling file type
      max_age: 7         # max expire days
      time_unit: day     # rolling time interval
`

type TestConfig struct {
	Log Config `yaml:"log"`
}

// func TestIllLogConfigPanic(t *testing.T) {
// 	var cfg TestConfig
// 	mustYamlUnmarshal(t, []byte(illConfigInfo), &cfg)
// 	conf := cfg.Log
// 	require.Panicsf(t, func() {
// 		plugin.Get("log", "default").Setup("default", &plugin.YamlNodeDecoder{Node: &conf})
// 	}, "NewRollWriter would return an error if file name is not configured")
// }

// type fakeDecoder struct{}

// func (c *fakeDecoder) Decode(conf interface{}) error {
// 	return nil
// }

// func mustYamlUnmarshal(t *testing.T, in []byte, out interface{}) {
// 	t.Helper()

// 	if err := yaml.Unmarshal(in, out); err != nil {
// 		t.Fatal(err)
// 	}
// }
