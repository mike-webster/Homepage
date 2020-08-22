package log

import (
	"context"
	"fmt"
	"time"

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
	k := ctx.Value(keys.LogLevel)
	s := ctx.Value(string(keys.LogLevel))

	ll := "debug"

	if k == nil && s == nil {
		fmt.Println("no log level in context, using default")
	} else if s == nil {
		ll = k.(string)
	} else {
		ll = s.(string)
	}

	switch ll {
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

	fmt.Println(">====\nTime: ", time.Now())
	for k, v := range payload {
		fmt.Printf("-- %v: \n\t%v\n", k, v)
	}
}
