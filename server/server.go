package server

import (
	"context"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	"homepage/log"
)

var (
	PathHomepage = "/"
)

func StartAPI(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Log(ctx, map[string]interface{}{
				"error": r,
				"event": "api_panic",
				"stack": string(debug.Stack()),
			}, "error")
		}
	}()
	runServer(ctx)
}

func runServer(ctx context.Context) {
	r := gin.New()

	r.LoadHTMLGlob("static/templates/*")

	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")

	r.GET(PathHomepage, func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", nil)
	})

	r.Run()
}
