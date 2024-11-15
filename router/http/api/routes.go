package api

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/internal/middleware/auth"
	v1 "github.com/terrapi-solution/controller/router/http/api/v1"
)

var token = "MIICmzCCAYMCBgGTEbyGnDANBgkqhkiG9w0BAQsFADARMQ8wDQYDVQQDDAZtYXN0ZXIwHhcNMjQxMTA5MTYyMTI4WhcNMzQxMTA5MTYyMzA4WjARMQ8wDQYDVQQDDAZtYXN0ZXIwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCu71rXnBXsGdkTqDsHdA/duTeSHy4sH4newOONeEKIsl/WVXb7PTtUtmyWk5yJYQYv18E1YPoUKp3l0w0bv8zWpOppixHwytgZRKxb3rK0FbGe/rjCcNFgM8QfJkyQa96ISkUaL0ljGf9O07p7PPb4Z4xvxpNsWcy2HOE4wr67Jc1qXhpbMmsCAaWznANHJHAj+qanwmS+Nu5BdWWKXdYB3r3feHqtgnSj/lLmR25ehPSAzrod0VlCDbEAwABoHtWGazIEnX34RuTiNI8gdcqZ43BUEjmwqPmyRVrfqJQoK4mKlsga1GcbYEP1HC9pU12oDenTI4zJX6xe94XLtWzlAgMBAAEwDQYJKoZIhvcNAQELBQADggEBAEJKT0BNzIgV8D3A98sawfktIeIbV5rr0xeWRI4TncOV5lZoN0jGO7w10NB5AgFQLejCZ0QzY0iAGQRoRhWozB5sO4u77xt938vYOkJ9U2h7fDMyiF2bhZVVGs5VDmpCJWT9sB3fT/C7f+0scygdz8kidB638oGdESbeEYfzRl9mNyQAbGhtY5FPwanHBywElGLRDU6lpX2PNc+e5oKNadf766GQMHPZx/OML9oofceEzW00A7KIkR+NiV5zhokkl8nd7y+QrIrvB24lzeCx3cq5Ts2SRHT2MmQdlcTcWwCACxGtYpn2OrCKqaqHhVpTl0f9G89nXlXWrbpMTOKrZQ0="
var SecretKey = "-----BEGIN CERTIFICATE-----\n" + token + "\n-----END CERTIFICATE-----"

// NewRoutesFactory creates and returns a factory to create routes for the v1 API
func NewRoutesFactory() func(group *gin.RouterGroup) {
	return func(group *gin.RouterGroup) {

		// Enable the auth middleware
		group.Use(auth.ValidateToken(SecretKey))

		// users route
		v1ModulesGroup := group.Group("/api")
		v1.NewRoutesFactory()(v1ModulesGroup)
	}
}
