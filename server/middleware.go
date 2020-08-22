package server

import (
	"homepage/env"
	"homepage/keys"
	"homepage/log"
	"os"

	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func loadContextValues(c *gin.Context) {
	cfg, err := env.GetConfig()
	if err != nil {
		panic(err)
	}
	logLevel := os.Getenv("LOG_LEVEL")
	if len(logLevel) < 1 {
		logLevel = "debug"
	}
	c.Set(string(keys.LogLevel), logLevel)

	sportsApiKey := os.Getenv("SPORTS_API")
	if len(logLevel) < 1 {
		panic("cant run without sports db api key")
	}

	c.Set(string(keys.SportsDB), sportsApiKey)
	c.Set(string(keys.AppConfig), cfg)
}

func recovery(c *gin.Context) {
	defer func() {
		type viewbag struct {
			Error map[string]interface{}
		}
		if r := recover(); r != nil {
			info := map[string]interface{}{
				"error": r,
				"event": "api_panic",
				"stack": string(debug.Stack()),
			}
			vb := viewbag{
				Error: info,
			}
			log.Log(c, info, "error")
			c.HTML(500, "err500.tmpl", vb)
		}
	}()
	c.Next()
}
