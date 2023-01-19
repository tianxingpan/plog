package plog

func init() {
	RegisterWriter(OutputConsole, DefaultConsoleWriterFactory)
	RegisterWriter(OutputFile, DefaultFileWriterFactory)
	RegisterLogger(defaultLoggerName, NewZapLog(defaultConfig))
}
