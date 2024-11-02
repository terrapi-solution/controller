package errors

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	// NotFound error indicates a missing / not found record
	NotFound        = "NotFound"
	notFoundMessage = "Record not found"

	// ValidationError indicates an error in input validation
	ValidationError        = "ValidationError"
	validationErrorMessage = "Validation error"

	// ResourceAlreadyExists indicates a duplicate / already existing record
	ResourceAlreadyExists     = "ResourceAlreadyExists"
	alreadyExistsErrorMessage = "Resource already exists"

	// RepositoryError indicates a repository (e.g database) error
	RepositoryError        = "RepositoryError"
	repositoryErrorMessage = "Error in repository operation"

	// NotAuthenticated indicates an authentication error
	NotAuthenticated             = "NotAuthenticated"
	notAuthenticatedErrorMessage = "Not Authenticated"

	// TokenGeneratorError indicates an token generation error
	TokenGeneratorError        = "TokenGeneratorError"
	tokenGeneratorErrorMessage = "Error in token generation"

	// NotAuthorized indicates an authorization error
	NotAuthorized             = "NotAuthorized"
	notAuthorizedErrorMessage = "Not Authorized"

	// UnknownError indicates an error that the app cannot find the cause for
	UnknownError        = "UnknownError"
	unknownErrorMessage = "Something went wrong"
)

// AppError defines an application (domain) error
type AppError struct {
	Err  error
	Type string
}

// NewAppError initializes a new domain error using an error and its type.
func NewAppError(err error, errType string) *AppError {
	return &AppError{
		Err:  err,
		Type: errType,
	}
}

// NewAppErrorWithType initializes a new default error for a given type.
func NewAppErrorWithType(errType string) *AppError {
	var err error

	switch errType {
	case NotFound:
		err = errors.New(notFoundMessage)
	case ValidationError:
		err = errors.New(validationErrorMessage)
	case ResourceAlreadyExists:
		err = errors.New(alreadyExistsErrorMessage)
	case RepositoryError:
		err = errors.New(repositoryErrorMessage)
	case NotAuthenticated:
		err = errors.New(notAuthenticatedErrorMessage)
	case NotAuthorized:
		err = errors.New(notAuthorizedErrorMessage)
	case TokenGeneratorError:
		err = errors.New(tokenGeneratorErrorMessage)
	default:
		err = errors.New(unknownErrorMessage)
	}

	return &AppError{
		Err:  err,
		Type: errType,
	}
}

// String converts the app error to a human-readable string.
func (appErr *AppError) Error() string {
	return appErr.Err.Error()
}

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
			var err *AppError
			ok := errors.As(errs[0].Err, &err)
			if ok {
				// Handle the error based on its type
				switch err.Type {
				case NotFound:
					c.JSON(http.StatusNotFound, err.Error())
					return
				case ValidationError:
					c.JSON(http.StatusBadRequest, err.Error())
					return
				case ResourceAlreadyExists:
					c.JSON(http.StatusConflict, err.Error())
					return
				case NotAuthenticated:
					c.JSON(http.StatusUnauthorized, err.Error())
					return
				case NotAuthorized:
					c.JSON(http.StatusForbidden, err.Error())
					return
				case RepositoryError:
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
