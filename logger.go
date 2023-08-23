package plog

// Field is the user defined log field.
type Field struct {
	Key   string
	Value interface{}
}

// Logger is the underlying logging work
type Logger interface {
	// Trace logs to TRACE log. Arguments are handled in the manner of fmt.Print.
	Trace(args ...interface{})
	// Tracef logs to TRACE log. Arguments are handled in the manner of fmt.Printf.
	Tracef(format string, args ...interface{})
	// Debug logs to DEBUG log. Arguments are handled in the manner of fmt.Print.
	Debug(args ...interface{})
	// Debugf logs to DEBUG log. Arguments are handled in the manner of fmt.Printf.
	Debugf(format string, args ...interface{})
	// Info logs to INFO log. Arguments are handled in the manner of fmt.Print.
	Info(args ...interface{})
	// Infof logs to INFO log. Arguments are handled in the manner of fmt.Printf.
	Infof(format string, args ...interface{})
	// Warn logs to WARNING log. Arguments are handled in the manner of fmt.Print.
	Warn(args ...interface{})
	// Warnf logs to WARNING log. Arguments are handled in the manner of fmt.Printf.
	Warnf(format string, args ...interface{})
	// Error logs to ERROR log. Arguments are handled in the manner of fmt.Print.
	Error(args ...interface{})
	// Errorf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
	Errorf(format string, args ...interface{})
	// Fatal logs to ERROR log. Arguments are handled in the manner of fmt.Print.
	// All Fatal logs will exit by calling os.Exit(1).
	// Implementations may also call os.Exit() with a non-zero exit code.
	Fatal(args ...interface{})
	// Fatalf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
	Fatalf(format string, args ...interface{})

	// Sync calls the underlying Core's Sync method, flushing any buffered log entries.
	// Applications should take care to call Sync before exiting.
	Sync() error

	// SetLevel set the output log level.
	SetLevel(output string, level Level)
	// GetLevel get the output log level.
	GetLevel(output string) Level
	// WithFields set some user defined data to logs, such as uid, imei, etc.
	// Fields must be paired.
	// Deprecated: use With instead.
	WithFields(fields ...string) Logger
	// With add user defined fields to Logger. Fields support multiple values.
	With(fields ...Field) Logger
}

// OptionLogger defines logger with additional options.
type OptionLogger interface {
	WithOptions(opts ...Option) Logger
}

// // DecoderImp decodes the log.
// type DecoderImp struct {
// 	OutputConfig *OutputConfig
// 	Core         zapcore.Core
// 	ZapLevel     zap.AtomicLevel
// }

// // Decode decodes writer configuration, copy one.
// func (d *DecoderImp) Decode(cfg interface{}) error {
// 	output, ok := cfg.(**OutputConfig)
// 	if !ok {
// 		return fmt.Errorf("decoder config type:%T invalid, not **OutputConfig", cfg)
// 	}
// 	*output = d.OutputConfig
// 	return nil
// }

// RegisterLogger registers Logger. It supports multiple Logger implementation.
func RegisterLogger(name string, logger Logger) {
	mu.Lock()
	defer mu.Unlock()
	if logger == nil {
		panic("log: Register logger is nil")
	}
	if _, dup := loggers[name]; dup && name != defaultLoggerName {
		panic("log: Register called twiced for logger name " + name)
	}
	loggers[name] = logger
	if name == defaultLoggerName {
		DefaultLogger = logger
	}
}

// GetLogger returns the Logger implementation by log name.
// log.Debug use DefaultLogger to print logs. You may also use log.Get("name").Debug.
func GetLogger(name string) Logger {
	mu.RLock()
	l := loggers[name]
	mu.RUnlock()
	return l
}
