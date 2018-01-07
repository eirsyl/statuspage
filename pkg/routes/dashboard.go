package routes

import (
	"github.com/eirsyl/statuspage/pkg"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
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

	owner, color, logo := viper.GetString("siteOwner"), viper.GetString("siteColor"), viper.GetString("siteLogo")

	if owner == "" {
		log.Fatal("The owner cannot be empty")
	}

	if color == "" {
		log.Fatal("The color cannot be empty")
	}

	if logo == "" {
		log.Fatal("The logo cannot be empty")
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
