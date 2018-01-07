package routes

import (
	"github.com/eirsyl/statuspage/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Dashboard(c *gin.Context) {
	services := c.Keys["services"].(pkg.Services)
	incidents := c.Keys["incidents"].(pkg.Incidents)

	res, err := services.GetServices()
	if err != nil {
		panic(err)
	}

	inc, err := incidents.GetLatestIncidents()
	if err != nil {
		panic(err)
	}

	owner := os.Getenv("SITE_OWNER")
	if owner == "" {
		owner = "Abakus"
	}

	color := os.Getenv("SITE_COLOR")
	if color == "" {
		color = "#343434"
	}

	logo := os.Getenv("SITE_LOGO")
	if logo == "" {
		logo = "static/img/logo.png"
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"owner":              owner,
		"backgroundColor":    color,
		"logo":               logo,
		"services":           pkg.AggregateServices(res),
		"mostCriticalStatus": pkg.MostCriticalStatus(res),
		"incidents":          pkg.AggregateIncidents(inc),
	})
}
