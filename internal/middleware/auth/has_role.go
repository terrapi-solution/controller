package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

// configuration for the HasRole middleware
type hasRoleConfig struct {
	JWTContextKey string
	RolesClaim    string
}

// Default configuration for the HasRole middleware
var defaultHasRoleConfig = hasRoleConfig{
	JWTContextKey: "token",
	RolesClaim:    "roles",
}

// HasRole middleware, depends on registration of auth.Validator
func HasRole(roles ...string) gin.HandlerFunc {
	return hasRoleWithConfig(defaultHasRoleConfig, roles...)
}

// hasRoleWithConfig is a middleware to check if the user has the required role
func hasRoleWithConfig(config hasRoleConfig, roles ...string) gin.HandlerFunc {
	// contains checks if a string is in a slice of strings
	var contains = func(e string, s []interface{}) bool {
		for _, n := range s {
			if e == n {
				return true
			}
		}
		return false
	}

	return func(c *gin.Context) {
		token, exists := c.Get(config.JWTContextKey)
		if !exists {
			respondWithError(c, http.StatusUnauthorized, "token is required")
			return
		}
		if token == nil {
			respondWithError(c, http.StatusUnauthorized, "user not found")
			return
		}

		claims := token.(*jwt.Token).Claims.(jwt.MapClaims)
		userRoles, ok := claims[config.RolesClaim].([]interface{})
		if !ok {
			respondWithError(c, http.StatusForbidden, "forbidden")
			return
		}

		for _, r := range roles {
			if !contains(r, userRoles) {
				respondWithError(c, http.StatusForbidden, fmt.Sprintf("forbidden: %s", r))
				return
			}
		}

		c.Next()
	}
}

// respondWithError sends a JSON response with the given status and message
func respondWithError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"status": status, "message": message})
	c.Abort()
}
