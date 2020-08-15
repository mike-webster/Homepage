package log

import (
	"context"
	"fmt"

	"github.com/mike-webster/homepage/keys"
)

var (
	// LogLevelDebug is for debug logs
	LogLevelDebug = "debug"
	// LogLevelInfo is for info logs
	LogLevelInfo = "info"
	// LogLevelError is for error logs
	LogLevelError = "error"
)

func Log(ctx context.Context, payload map[string]interface{}, level string) {
	ll := ctx.Value(keys.LogLevel)
	if ll == nil {
		return
	}

	strLL, ok := ll.(string)
	if !ok {
		return
	}

	switch strLL {
	case LogLevelError:
		if !(level == LogLevelError) {
			return
		}
	case LogLevelInfo:
		if !(level == LogLevelInfo || level == LogLevelError) {
			return
		}
	case LogLevelDebug:
	}

	fmt.Println()
}
