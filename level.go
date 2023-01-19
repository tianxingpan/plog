package plog

// Level is the log level.
type Level int

// Enums log level constants.
const (
	LevelNil Level = iota
	LevelTrace
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

// String turns the LogLevel to string.
func (lv *Level) String() string {
	return LevelStrings[*lv]
}

// LevelStrings is the map from log level to its string representation.
var LevelStrings = map[Level]string{
	LevelTrace: "trace",
	LevelDebug: "debug",
	LevelInfo:  "info",
	LevelWarn:  "warn",
	LevelError: "error",
	LevelFatal: "fatal",
}

// LevelNames is the map from string to log level.
var LevelNames = map[string]Level{
	"trace": LevelTrace,
	"debug": LevelDebug,
	"info":  LevelInfo,
	"warn":  LevelWarn,
	"error": LevelError,
	"fatal": LevelFatal,
}
