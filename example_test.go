package plog_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	log "github.com/tianxingpan/plog"
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
