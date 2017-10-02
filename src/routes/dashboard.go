package routes

import (
	"github.com/eirsyl/statuspage/src"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Dashboard(c *gin.Context) {
	services := c.Keys["services"].(src.Services)
	incidents := c.Keys["incidents"].(src.Incidents)

	res, err := services.GetServices()
	if err != nil {
		panic(err)
	}

	inc, err := incidents.GetLatestIncidents()
	if err != nil {
		panic(err)
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"owner":              "Abakus",
		"services":           src.AggregateServices(res),
		"mostCriticalStatus": src.MostCriticalStatus(res),
		"incidents":          src.AggregateIncidents(inc),
	})
}
