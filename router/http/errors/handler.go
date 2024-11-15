package errors

import "github.com/gin-gonic/gin"

// HandlerWithError is a function that handles an HTTP request.
// This is essentially the same as [gin.HandlerFunc], but allows returning an error.
type HandlerWithError func(c *gin.Context) error

// HandlerWithErrorWrapper wraps a [HandlerWithError] into a [gin.HandlerFunc].
// If an error is returned, it adds it to the context with c.Error(err).
// This needs to be handled with a middleware, such as [ErrorMiddleware]
func HandlerWithErrorWrapper(h HandlerWithError) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := h(c); err != nil {
			c.Abort()
			_ = c.Error(err)
		}
	}
}
