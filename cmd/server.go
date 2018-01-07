package cmd

import (
	"github.com/eirsyl/statuspage/pkg"
	"github.com/eirsyl/statuspage/pkg/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	serverCmd.PersistentFlags().StringP("listenAddress", "l", ":8080", "The address the server should listen on")
	serverCmd.PersistentFlags().StringP("token", "t", "", "The token used for authorizing with the API")
	serverCmd.PersistentFlags().StringP("postgresAddress", "a", "127.0.0.1:5432", "The address postgres is listening on")
	serverCmd.PersistentFlags().StringP("postgresUser", "u", "statuspage", "The user statuspage should connect to postgres as")
	serverCmd.PersistentFlags().StringP("postgresPassword", "p", "", "The postgres password")
	serverCmd.PersistentFlags().StringP("postgresDatabase", "d", "statuspage", "The postgres database statuspage should use")
	serverCmd.PersistentFlags().StringP("siteOwner", "", "Statuspage", "Site owner, used by the html templates")
	serverCmd.PersistentFlags().StringP("siteColor", "", "#343434", "The top color used in the templates")
	serverCmd.PersistentFlags().StringP("siteLogo", "", "static/img/logo.png", "Path to logo")
	viper.BindPFlag("listenAddress", serverCmd.PersistentFlags().Lookup("listenAddress"))
	viper.BindPFlag("token", serverCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("postgresAddress", serverCmd.PersistentFlags().Lookup("postgresAddress"))
	viper.BindPFlag("postgresUser", serverCmd.PersistentFlags().Lookup("postgresUser"))
	viper.BindPFlag("postgresPassword", serverCmd.PersistentFlags().Lookup("postgresPassword"))
	viper.BindPFlag("postgresDatabase", serverCmd.PersistentFlags().Lookup("postgresDatabase"))
	viper.BindPFlag("siteOwner", serverCmd.PersistentFlags().Lookup("siteOwner"))
	viper.BindPFlag("siteColor", serverCmd.PersistentFlags().Lookup("siteColor"))
	viper.BindPFlag("siteLogo", serverCmd.PersistentFlags().Lookup("siteLogo"))
	viper.BindEnv("listenAddress", "LISTEN_ADDRESS")
	viper.BindEnv("token", "TOKEN")
	viper.BindEnv("postgresAddress", "POSTGRES_ADDRESS")
	viper.BindEnv("postgresUser", "POSTGRES_USER")
	viper.BindEnv("postgresPassword", "POSTGRES_PASSWORD")
	viper.BindEnv("postgresDatabase", "POSTGRES_DATABASE")
	viper.BindEnv("siteOwner", "SITE_OWNER")
	viper.BindEnv("siteColor", "SITE_COLOR")
	viper.BindEnv("siteLogo", "SITE_LOGO")
	RootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the statuspage server",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.ConfigRuntime()
		gin.SetMode(gin.ReleaseMode)

		router := gin.New()
		router.Use(gin.Recovery())
		router.Use(pkg.State())
		router.Use(pkg.Logger())

		binding.Validator.RegisterValidation("incidentstatus", pkg.IncidentStatus)
		binding.Validator.RegisterValidation("servicestatus", pkg.ServiceStatus)

		router.Static("/static", "./static")
		router.LoadHTMLGlob("templates/*")

		router.GET("/", routes.Dashboard)

		api := router.Group("/api")
		api.Use(pkg.Auth())
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

		listenAddress := viper.GetString("listenAddress")
		log.WithFields(log.Fields{"address": listenAddress}).Info("Starting server")
		err := router.Run(listenAddress)
		if err != nil {
			log.Error("Server exited ", err)
		}
	},
}
