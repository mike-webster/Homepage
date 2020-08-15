package main

import (
	"context"

	"github.com/mike-webster/homepage/keys"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, keys.LogLevel, "debug")
	ctx = context.WithValue(ctx, keys.SportsDB, 1)
}
