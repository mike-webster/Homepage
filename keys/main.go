package keys

// ContextKey is the datatype to use to store/retrieve data in the context
type ContextKey string

var (
	// LogLevel is the key to use for the log level
	LogLevel ContextKey = "log-level"
	// SportsDB is the key to use for the sportsdb api
	SportsDB ContextKey = "sports-db"
	// AppConfig is the key to get the config
	AppConfig ContextKey = "app-config"
)
