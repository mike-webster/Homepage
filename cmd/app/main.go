package main

import (
	"context"
	"os"

	"homepage/keys"
	"homepage/server"
)

var (
	envLogLevel = "LOG_LEVEL"
	envSportsDB = "SPORTS_API"
)

func main() {
	ctx := readEnvironmentVariables()
	server.StartAPI(ctx)
}

func readEnvironmentVariables() context.Context {
	ctx := context.Background()
	logLevel := os.Getenv(envLogLevel)
	if len(logLevel) < 1 {
		logLevel = "debug"
	}
	ctx = context.WithValue(ctx, keys.LogLevel, logLevel)

	sportsApiKey := os.Getenv(envSportsDB)
	if len(logLevel) < 1 {
		panic("cant run without sports db api key")
	}
	return context.WithValue(ctx, keys.SportsDB, sportsApiKey)
}
