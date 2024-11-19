package api

import (
	"github.com/gin-gonic/gin"
	"github.com/terrapi-solution/controller/internal/middleware/auth"
	v1 "github.com/terrapi-solution/controller/router/http/api/v1"
)

var token = "MIICmzCCAYMCBgGTRJloIjANBgkqhkiG9w0BAQsFADARMQ8wDQYDVQQDDAZtYXN0ZXIwHhcNMjQxMTE5MTMyMzQ0WhcNMzQxMTE5MTMyNTI0WjARMQ8wDQYDVQQDDAZtYXN0ZXIwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCukOk4I53lEgas7zDOwnZ1W6EUwPAPnAeruonF2KcIZKH8drMS7Z3BJobFOEwMrFKCswAH59rwXPYT53wIH8SdvXj2++1W4v/5dMQ8Td/ZZyyL8ApeQtu90jn/TWKXBDEIYpiZsbN+wnfSHVIil8XISKsvJcLfw7iJ8mSnc2Rvz6S7Lkn2TgK+roKYXlGMTBKnF8Ic0vgzfDp4kpSMR4kyQfvgIgATHRHRpZf1wOzAbU473SZCKj5IFpEFzs+s+SUS4lytHC3X+bnm9zPnKbGqzRpwyHTHYGJ49/b1QApcSK1Ie2TZag9BuI+Q6Uvllq5R0W9DcICGwIIExvZNXLdBAgMBAAEwDQYJKoZIhvcNAQELBQADggEBAERD1DBWFF7bb97q1944O44iRI0w9+N9m2by/QQdBH6pZowgt6nqNmsVrq2hzz5kgzllCfAUs+zM+ue88bdq3CahdEQgZzYGk3BfF3leSLljIWsS7OTsu/xuglcMfodmhI9zEaGfVdcUbx397/mbonXD+aUU/qedVVj/ybTWWloOuhk85UEjn5413bsXgHZcXCSTxnMLsOKE29NM7fTOTqxKYcf9mM44/+Xn9hWnFjAOz4xsR9mjRAM3wDQJn99Gm044y/6Gvxt41Sg3xKrAmpX8OK7HWw5pGiokM5JP99x+amLT72clnd8awBtVhD/fBZ3VLIP6LAiJHp+d4j0hkTY="
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
