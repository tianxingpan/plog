package plog_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	log "github.com/tianxingpan/plog"
	"gopkg.in/yaml.v3"
)

func TestSetLevel(t *testing.T) {
	const level = "0"
	log.SetLevel(level, log.LevelInfo)
	require.Equal(t, log.LevelInfo, log.GetLevel(level))
}

func TestSetLogger(t *testing.T) {
	logger := log.NewZapLog(log.Config{})
	log.SetLogger(logger)
	require.Equal(t, log.GetDefaultLogger(), logger)
}

func TestLogXXX(t *testing.T) {
	log.Fatal("xxx")
}

func Example() {
	l := log.NewZapLog([]log.OutputConfig{
		{
			Writer:       "console",
			Level:        "debug",
			Formatter:    "console",
			FormatConfig: log.FormatConfig{TimeFmt: "xxx"},
		},
	})
	const defaultLoggerName = "default"
	oldDefaultLogger := log.GetDefaultLogger()
	log.Register(defaultLoggerName, l)
	defer func() {
		log.Register(defaultLoggerName, oldDefaultLogger)
	}()

	l = l.With(log.Field{Key: "plog", Value: "log"})
	l.Trace("hello world")
	l.Debug("hello world")
	l.Info("hello world")
	l.Warn("hello world")
	l.Error("hello world")
	l.Tracef("hello world")
	l.Debugf("hello world")
	l.Infof("hello world")
	l.Warnf("hello world")
	l.Errorf("hello world")

	// Output:
	// xxx	DEBUG	log/example_test.go:22	hello world	{"plog": "log"}
	// xxx	DEBUG	log/example_test.go:23	hello world	{"plog": "log"}
	// xxx	INFO	log/example_test.go:24	hello world	{"plog": "log"}
	// xxx	WARN	log/example_test.go:25	hello world	{"plog": "log"}
	// xxx	ERROR	log/example_test.go:26	hello world	{"plog": "log"}
	// xxx	DEBUG	log/example_test.go:27	hello world	{"plog": "log"}
	// xxx	DEBUG	log/example_test.go:28	hello world	{"plog": "log"}
	// xxx	INFO	log/example_test.go:29	hello world	{"plog": "log"}
	// xxx	WARN	log/example_test.go:30	hello world	{"plog": "log"}
	// xxx	ERROR	log/example_test.go:31	hello world	{"plog": "log"}
}

func TestLogFactory(t *testing.T) {

	log.EnableTrace()

	f := &log.Factory{}

	assert.Equal(t, "log", f.Type())

	// empty decoder
	err := f.Setup("default", nil)
	assert.NotNil(t, err)

	log.Register("default", log.DefaultLogger)
	assert.Equal(t, log.DefaultLogger, log.Get("default"))
	assert.Nil(t, log.Get("empty"))
	log.Sync()

	logger := log.WithFields("uid", "1111")
	assert.NotNil(t, logger)
	logger.Debugf("test")

	log.Trace("test")
	log.Tracef("test %s", "s")
	log.Debug("test")
	log.Debugf("test %s", "s")
	log.Error("test")
	log.Errorf("test %s", "s")
	log.Info("test")
	log.Infof("test %s", "s")
	log.Warn("test")
	log.Warnf("test %s", "s")
	log.Fatal("test %s", "s")
	log.Fatalf("test %s", "s")
}

func TestWriterFactory(t *testing.T) {
	t.Run("console", func(t *testing.T) {
		f := &log.ConsoleWriterFactory{}
		require.Equal(t, "log", f.Type())

		err := f.Setup("default", nil)
		require.Contains(t, err.Error(), "decoder empty")
	})
	t.Run("file", func(t *testing.T) {
		f := &log.FileWriterFactory{}
		require.Equal(t, "log", f.Type())

		err := f.Setup("default", nil)
		require.Contains(t, err.Error(), "decoder empty")
	})

}

const configInfo = `
log:
  - writer: console # default as console std output
    level: debug # std log level
  - writer: file # local log file
    level: debug # std log level
    writer_config: # config of local file output
      filename: trpc_time.log # the path of local rolling log files
      roll_type: time    # file rolling type
      max_age: 7         # max expire days
      time_unit: day     # rolling time interval
  - writer: file # local file log
    level: debug # std output log level
    writer_config: # config of local file output
      filename: trpc_size.log # the path of local rolling log files
      roll_type: size    # file rolling type
      max_age: 7         # max expire days
      max_size: 100      # size of local rolling file, unit MB
      max_backups: 10    # max number of log files
      compress:  false   # should compress log file
`

type TestConfig struct {
	Log log.Config `yaml:"log"`
}

func TestLogFactorySetup(t *testing.T) {
	oldDefaultLogger := log.GetDefaultLogger()
	defer func() {
		log.Register("default", oldDefaultLogger)
	}()

	var cfg TestConfig
	mustYamlUnmarshal(t, []byte(configInfo), &cfg)
	conf := cfg.Log
	err := log.Init(conf)
	assert.Nil(t, err)

	log.Trace("test")
	log.Debug("test")
	log.Error("test")
	log.Info("test")
	log.Warn("test")
}

func mustYamlUnmarshal(t *testing.T, in []byte, out interface{}) {
	t.Helper()

	if err := yaml.Unmarshal(in, out); err != nil {
		t.Fatal(err)
	}
}
