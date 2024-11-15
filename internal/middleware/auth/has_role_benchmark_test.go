package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkHasRole_WithoutToken(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(HasRole("admin"))
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})

	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
	}
}

func BenchmarkHasRole_WithoutTokenHasAdminRole(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	tokenString := createRoleToken([]interface{}{"user"})

	r := gin.New()
	r.Use(setTokenMiddleware(tokenString))
	r.Use(HasRole("admin"))
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})

	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
	}
}

func BenchmarkHasRole_WithTokenHasAdminRole(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	tokenString := createRoleToken([]interface{}{"admin"})

	r := gin.New()
	r.Use(setTokenMiddleware(tokenString))
	r.Use(HasRole("admin"))
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})

	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
	}
}

func createRoleToken(roles []interface{}) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"roles": roles})
	tokenString, _ := token.SignedString([]byte("secret"))
	return tokenString
}

func setTokenMiddleware(tokenString string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		c.Set("token", token)
		c.Next()
	}
}
