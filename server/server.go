package server

import (
	"context"
	"fmt"
	"net/http"
	"sort"

	"homepage/data"
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
		cfg, err := env.GetConfigFromContext(c)
		if err != nil {
			panic(err)
		}

		type viewbag struct {
			Events data.Events
		}

		info := viewbag{}

		for _, i := range cfg.Teams {
			log.Log(ctx, map[string]interface{}{"event": "retrieving_events", "team": i}, "info")
			t, err := data.GetTeamByName(ctx, i)
			if err != nil {
				panic(err)
			}
			es, err := data.GetNext5Events(ctx, t.ID)
			if err != nil {
				panic(err)
			}
			info.Events = append(info.Events, *es...)
		}

		sort.Sort(info.Events)

		c.HTML(http.StatusOK, "home.tmpl", info)
	})

	r.Run(fmt.Sprintf("%v:%v", "127.0.0.1", cfg.Port))
}
