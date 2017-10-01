package main

import (
	"github.com/gin-gonic/gin"
	"github.com/eirsyl/statuspage/src"
	"github.com/eirsyl/statuspage/src/routes"
	"os"
	"github.com/go-pg/pg"
	"runtime"
	"log"
)

func main()  {

	ConfigRuntime()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(State())

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", routes.Dashboard)

	router.GET("/api/services", routes.ServiceList)
	router.POST("/api/services", routes.ServicePost)
	router.GET("/api/services/:id", routes.ServiceGet)
	router.PATCH("/api/services/:id", routes.ServicePatch)
	router.DELETE("/api/services/:id", routes.ServiceDelete)

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
		Addr: pgAddr,
		User: pgUser,
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