package errors

import (
	"errors"
	"github.com/gin-gonic/gin"
	domain "github.com/terrapi-solution/controller/domain/errors"
	"net/http"
)

// Handler is Gin middleware to handle error.
func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Execute request handlers and then handle any error
		c.Next()

		// Retrieve the list of error from the context
		errs := c.Errors

		// If there are any error, handle the first one
		if len(errs) > 0 {
			// Attempt to cast the error to an AppError
			var err *domain.AppError
			ok := errors.As(errs[0].Err, &err)
			if ok {
				// Handle the error based on its type
				switch err.Type {
				case domain.NotFound:
					c.JSON(http.StatusNotFound, err.Error())
					return
				case domain.ValidationError:
					c.JSON(http.StatusBadRequest, err.Error())
					return
				case domain.ResourceAlreadyExists:
					c.JSON(http.StatusConflict, err.Error())
					return
				case domain.NotAuthenticated:
					c.JSON(http.StatusUnauthorized, err.Error())
					return
				case domain.NotAuthorized:
					c.JSON(http.StatusForbidden, err.Error())
					return
				case domain.RepositoryError:
					// Fall through to the default case
				default:
					c.JSON(http.StatusInternalServerError, err.Error())
					return
				}
			}

			// If the error is not an AppError, return a generic internal server error
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}
	}
}
