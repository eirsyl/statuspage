package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

// Auth enforces token authentication on the api endpoints.
func Auth() gin.HandlerFunc {
	validToken := viper.GetString("token")

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if validToken == "" || token != validToken {
			log.WithFields(log.Fields{"token": token}).Warn("Permission denied, invalid token")
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

// Logger add logrus logging to each request
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

		if raw != "" {
			path = path + "?" + raw
		}

		log.WithFields(log.Fields{
			"statusCode": statusCode,
			"latency":    latency,
			"clientIP":   clientIP,
			"method":     method,
			"path":       path,
			"comment":    comment,
		}).Info("Request")

	}
}
