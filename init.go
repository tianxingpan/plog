package plog

import "github.com/tianxingpan/plog/plugin"

func init() {
	RegisterWriter(OutputConsole, DefaultConsoleWriterFactory)
	RegisterWriter(OutputFile, DefaultFileWriterFactory)
	RegisterLogger(defaultLoggerName, NewZapLog(defaultConfig))
	plugin.Register(defaultLoggerName, DefaultLogFactory)
}
