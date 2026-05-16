package logger

// Writer is an object that provides a log method.
type Writer interface {
	Log(Level, string, ...any)
}

// FieldWriter is a Writer that also supports structured key-value field logging.
// Logger implements this interface. Components that want structured output
// (e.g. path lifecycle events with trace_id) should accept FieldWriter.
type FieldWriter interface {
	Writer
	LogFields(level Level, msg string, fields map[string]string)
}
