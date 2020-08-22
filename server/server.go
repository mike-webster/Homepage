package server

import (
	"context"
	"fmt"
	"net/http"

	"homepage/env"
	"homepage/log"

	"github.com/gin-gonic/gin"
)

var (
	PathHomepage = "/"
)

func StartAPI(ctx context.Context) {
	runServer(ctx)
}

func runServer(ctx context.Context) {
	cfg, err := env.GetConfigFromContext(ctx)
	if err != nil {
		panic(err)
	}
	r := gin.New()

	r.Use(recovery)
	r.Use(loadContextValues)

	r.LoadHTMLGlob("static/templates/*")

	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")

	r.GET(PathHomepage, func(c *gin.Context) {
		log.Log(c, map[string]interface{}{"event": "incoming"}, "info")
		cfg, err := env.GetConfigFromContext(c)
		if err != nil {
			panic(err)
		}

		type viewbag struct {
			Teams   []string
			Leagues []string
		}

		vb := viewbag{
			Teams:   cfg.Teams,
			Leagues: cfg.Leagues,
		}

		// LEFT OFF: this isn't working - showing an empty slice on the page
		c.HTML(http.StatusOK, "home.tmpl", vb)
	})

	r.Run(fmt.Sprintf("%v:%v", "127.0.0.1", cfg.Port))
}
