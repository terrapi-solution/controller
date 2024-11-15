package errors

import (
	"github.com/gin-gonic/gin"
	domainErrors "github.com/terrapi-solution/controller/domain/errors"
	"net/http"
)

// Handler returns a gin middleware that handles errors.
func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		handleError(c)
	}
}

// handleError returns a gin middleware which writes a response with the error in the context.
func handleError(c *gin.Context) {
	// Get the last error and return if there is none
	err := c.Errors.Last()
	if err == nil {
		return
	}

	// Convert the error to a domain error
	appErr := domainErrors.ToError(err.Err)
	if appErr == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Write the error response if not already written
	if !c.Writer.Written() {
		c.JSON(appErr.HTTPStatusCode(), appErr)
	}
}
