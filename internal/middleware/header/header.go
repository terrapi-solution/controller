package header

import (
	"github.com/terrapi-solution/controller/internal/version"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Cache writes required cache headers to all requests.
func Cache() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
		c.Writer.Header().Set("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
		c.Writer.Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))

		c.Next()
	}
}

// Options writes required option headers to all requests.
func Options() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
			c.Writer.Header().Set("Allow", "HEAD, GET, POST, PUT, PATCH, DELETE, OPTIONS")

			c.Status(http.StatusOK)
			c.Writer.WriteHeaderNow()
			c.Abort()
		}
	}
}

// Secure writes required access headers to all requests.
func Secure() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")

		if c.Request.TLS != nil {
			c.Header("Strict-Transport-Security", "max-age=31536000")
		}

		c.Next()
	}
}

// Version writes the current API version to the headers.
func Version() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-TERRAPI-VERSION", version.String)

		c.Next()
	}
}
