package keys

// ContextKey is the datatype to use to store/retrieve data in the context
type ContextKey string

var (
	// LogLevel is the key to use to store/retrieve the log level from the context
	LogLevel ContextKey = "log-level"
)
