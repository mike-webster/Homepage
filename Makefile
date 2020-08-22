SPORTS_DB_KEY = 1
LOG = "debug"

.PHONY: start

start:
	SPORTS_API=$(SPORTS_DB_KEY) LOG_LEVEL=$(LOG) go run cmd/app/main.go