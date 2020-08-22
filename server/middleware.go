package server

import (
	"homepage/env"
	"homepage/keys"
	"homepage/log"

	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func loadContextValues(c *gin.Context) {
	cfg, err := env.GetConfig()
	if err != nil {
		panic(err)
	}
	c.Set(string(keys.AppConfig), cfg)
}

func recovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Log(c, map[string]interface{}{
				"error": r,
				"event": "api_panic",
				"stack": string(debug.Stack()),
			}, "error")
		}
	}()
	c.Next()
}
