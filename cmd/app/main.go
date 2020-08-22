package main

import (
	"context"
	"os"

	"homepage/env"
	"homepage/keys"
	"homepage/server"
)

var (
	envLogLevel = "LOG_LEVEL"
	envSportsDB = "SPORTS_API"
	configPath  = "app.yaml"
)

func main() {
	ctx := readEnvironmentVariables()
	server.StartAPI(ctx)
}

func readEnvironmentVariables() context.Context {
	ctx := context.Background()
	cfg, err := env.GetConfig()
	if err != nil {
		panic(err)
	}

	logLevel := os.Getenv(envLogLevel)
	if len(logLevel) < 1 {
		logLevel = "debug"
	}
	ctx = context.WithValue(ctx, keys.LogLevel, logLevel)

	sportsApiKey := os.Getenv(envSportsDB)
	if len(logLevel) < 1 {
		panic("cant run without sports db api key")
	}

	ctx = context.WithValue(ctx, keys.SportsDB, sportsApiKey)
	return context.WithValue(ctx, keys.AppConfig, cfg)
}
