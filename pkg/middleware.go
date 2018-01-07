package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/spf13/viper"
	"net/http"
)

// Auth enforces token authentication on the api endpoints.
func Auth() gin.HandlerFunc {
	validToken := viper.GetString("token")

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if validToken == "" || token != validToken {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

// State provides the handler with access to the datastore
func State() gin.HandlerFunc {
	pgUser, pgPassword := viper.GetString("postgresUser"), viper.GetString("postgresPassword")
	pgAddress, pgDatabase := viper.GetString("postgresAddress"), viper.GetString("postgresDatabase")

	db := pg.Connect(&pg.Options{
		Addr:     pgAddress,
		User:     pgUser,
		Password: pgPassword,
		Database: pgDatabase,
	})

	if err := CreateSchema(db); err != nil {
		panic(err)
	}

	services := Services{}
	services.Initialize(*db)

	incidents := Incidents{}
	incidents.Initialize(*db)

	return func(c *gin.Context) {
		c.Set("services", services)
		c.Set("incidents", incidents)
		c.Next()
	}
}
