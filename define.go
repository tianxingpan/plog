package plog

import "sync"

const (
	pluginType        = "log"
	defaultLoggerName = "default"
)

var (
	// DefaultLogger the default Logger. The initial output is console. When frame start, it is
	// over write by configuration.
	DefaultLogger Logger
	// DefaultLogFactory is the default log loader. Users may replace it with their own

	mu      sync.RWMutex
	loggers = make(map[string]Logger)
)
