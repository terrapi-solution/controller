package metric

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// Handler initializes the prometheus middleware.
func (s *MetricServer) metricHandler(token string) gin.HandlerFunc {
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
