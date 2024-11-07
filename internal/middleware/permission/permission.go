package permission

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/internal/user"
	"net/http"
)

// Handler is a middleware to check if the user has the required role
func Handler(roles ...user.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, err := getUserFromContext(c)
		if err != nil {
			respondWithError(c, http.StatusUnauthorized, "user not found")
			return
		}

		if !hasRequiredRole(u, roles) {
			respondWithError(c, http.StatusForbidden, "forbidden")
			return
		}

		c.Next()
	}
}

// respondWithError sends a JSON response with the given status and message
func respondWithError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"status": status, "message": message})
	c.Abort()
}

// hasRequiredRole checks if the user has the required role
func hasRequiredRole(user user.User, roles []user.Role) bool {
	for _, r := range roles {
		if user.Role == r {
			return true
		}
	}
	return false
}

// getUserFromContext is a helper function to get the user from the gin context
func getUserFromContext(c *gin.Context) (user.User, error) {
	// Retrieve the user from the context
	u, exists := c.Get("user")
	if !exists {
		return user.User{}, errors.New("user not found in the context")
	}

	// Parse the user from the context
	userParsed, valid := u.(*user.User)
	if !valid {
		return user.User{}, errors.New("unable to parse user from context")
	}

	return *userParsed, nil
}
