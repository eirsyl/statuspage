package cmd

import (
	"github.com/eirsyl/statuspage/pkg"
	"github.com/eirsyl/statuspage/pkg/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-pg/pg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

func init() {
	serverCmd.PersistentFlags().StringP("token", "t", "", "The token used for authorizing with the API")
	serverCmd.PersistentFlags().StringP("postgresAddress", "a", "statuspage", "The address postgres is listening on")
	serverCmd.PersistentFlags().StringP("postgresUser", "u", "statuspage", "The user statuspage should connect to postgres as")
	serverCmd.PersistentFlags().StringP("postgresPassword", "p", "", "The postgres password")
	serverCmd.PersistentFlags().StringP("postgresDatabase", "d", "statuspage", "The postgres database statuspage should use")
	viper.BindPFlag("token", serverCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("postgresAddress", serverCmd.PersistentFlags().Lookup("postgresAddress"))
	viper.BindPFlag("postgresUser", serverCmd.PersistentFlags().Lookup("postgresUser"))
	viper.BindPFlag("postgresPassword", serverCmd.PersistentFlags().Lookup("postgresPassword"))
	viper.BindPFlag("postgresDatabase", serverCmd.PersistentFlags().Lookup("postgresDatabase"))
	viper.BindEnv("token", "TOKEN")
	viper.BindEnv("postgresAddress", "POSTGRES_ADDRESS")
	viper.BindEnv("postgresUser", "POSTGRES_USER")
	viper.BindEnv("postgresPassword", "POSTGRES_PASSWORD")
	viper.BindEnv("postgresDatabase", "POSTGRES_DATABASE")
	RootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the statuspage server",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.ConfigRuntime()
		gin.SetMode(gin.ReleaseMode)

		router := gin.Default()
		router.Use(State())

		binding.Validator.RegisterValidation("incidentstatus", pkg.IncidentStatus)
		binding.Validator.RegisterValidation("servicestatus", pkg.ServiceStatus)

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
	},
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

	if err := pkg.CreateSchema(db); err != nil {
		panic(err)
	}

	services := pkg.Services{}
	services.Initialize(*db)

	incidents := pkg.Incidents{}
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
