package keys

// ContextKey is the datatype to use to store/retrieve data in the context
type ContextKey string

var (
	// LogLevel is the key to use for the log level
	LogLevel ContextKey = "log-level"
	// SportsDB is the key to use for the sports db api key
	SportsDB ContextKey = "sports-db"
)
