package server

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/terrapi-solution/controller/internal/config"
	"net/http"
)

// Metrics initializes the routing of metrics and health.
func Metrics(cfg *config.Config) *gin.Engine {
	// Creates a router without any middleware
	r := gin.New()

	// endpoint for prometheus metrics
	r.GET("/metrics", metricHandler(cfg.Metrics.Token))

	return r
}

// Handler initializes the prometheus middleware.
func metricHandler(token string) gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		if token == "" {
			h.ServeHTTP(c.Writer, c.Request)
			return
		}

		header := c.Request.Header.Get("Authorization")

		if header == "" {
			c.Status(http.StatusUnauthorized)
			c.Writer.WriteHeaderNow()
			c.Abort()
			return
		}

		if header != "Bearer "+token {
			c.Status(http.StatusUnauthorized)
			c.Writer.WriteHeaderNow()
			c.Abort()
			return
		}

		h.ServeHTTP(c.Writer, c.Request)
	}
}
