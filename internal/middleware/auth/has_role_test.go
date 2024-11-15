package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHasRole_WithoutToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new gin router
	r := gin.New()
	r.Use(HasRole("admin"))
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})

	// Create a new HTTP request
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "{\"message\":\"token is required\",\"status\":401}", w.Body.String())
}

func TestHasRole_WithoutTokenHasAdminRole(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a test token with "admin" role
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"roles": []interface{}{"user"},
	})
	tokenString, _ := token.SignedString([]byte("secret"))

	// Create a new gin router
	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set("token", token)
		c.Next()
	})
	r.Use(HasRole("admin"))
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})

	// Create a new HTTP request
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Equal(t, "{\"message\":\"forbidden: admin\",\"status\":403}", w.Body.String())
}

func TestHasRole_WithTokenHasAdminRole(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a test token with "admin" role
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"roles": []interface{}{"admin"},
	})
	tokenString, _ := token.SignedString([]byte("secret"))

	// Create a new gin router
	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set("token", token)
		c.Next()
	})
	r.Use(HasRole("admin"))
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})

	// Create a new HTTP request
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "success", w.Body.String())
}
