package main

import (
	"github.com/eirsyl/statuspage/src"
	"github.com/eirsyl/statuspage/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-pg/pg"
	"log"
	"net/http"
	"os"
	"runtime"
)

func main() {

	ConfigRuntime()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(State())

	binding.Validator.RegisterValidation("incidentstatus", src.IncidentStatus)
	binding.Validator.RegisterValidation("servicestatus", src.ServiceStatus)

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", routes.Dashboard)

	api := router.Group("/api")
	api.Use(Auth())
	{
		api.GET("/services", routes.ServiceList)
		api.POST("/services", routes.ServicePost)
		api.GET("/services/:id", routes.ServiceGet)
		api.PATCH("/services/:id", routes.ServicePatch)
		api.DELETE("/services/:id", routes.ServiceDelete)

		api.GET("/incidents", routes.IncidentList)
		api.POST("/incidents", routes.IncidentPost)
		api.GET("/incidents/:id", routes.IncidentGet)
		api.DELETE("/incidents/:id", routes.IncidentDelete)

		api.GET("/incidents/:id/updates", routes.IncidentUpdateList)
		api.POST("/incidents/:id/updates", routes.IncidentUpdatePost)
		api.GET("/incidents/:id/updates/:updateId", routes.IncidentUpdateGet)
		api.DELETE("/incidents/:id/updates/:updateId", routes.IncidentUpdateDelete)
	}

	router.Run()

}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	log.Printf("Running with %d CPUs\n", nuCPU)
}

func State() gin.HandlerFunc {
	pgAddr := os.Getenv("POSTGRES_ADDRESS")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgDB := os.Getenv("POSTGRES_DB")

	db := pg.Connect(&pg.Options{
		Addr:     pgAddr,
		User:     pgUser,
		Password: pgPassword,
		Database: pgDB,
	})

	if err := src.CreateSchema(db); err != nil {
		panic(err)
	}

	services := src.Services{}
	services.Initialize(*db)

	incidents := src.Incidents{}
	incidents.Initialize(*db)

	return func(c *gin.Context) {
		c.Set("services", services)
		c.Set("incidents", incidents)
		c.Next()
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		validToken := os.Getenv("API_TOKEN")

		if !(len(validToken) > 0 && token == validToken) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
