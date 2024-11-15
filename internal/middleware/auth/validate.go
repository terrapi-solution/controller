package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

// configuration for the Validate middleware
type validateConfig struct {
	publicKey *rsa.PublicKey
}

// ValidateToken is a middleware to validate the JWT token sent by the client
func ValidateToken(certificate string) gin.HandlerFunc {
	// Parse the public key from the certificate
	publicKey, err := parsePublicKey(certificate)
	if err != nil {
		panic(err)
	}

	// Return the validation middleware
	return validateWithConfig(validateConfig{publicKey})
}

// validateWithConfig is a middleware to validate the JWT token sent by the client
func validateWithConfig(config validateConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the user JWT from the Authorization header
		rawToken := extractTokenFromContext(c)

		// Validate the token sent by the client
		status, token := parseToken(c, rawToken, config.publicKey)
		if !status {
			return
		}

		// Store the information in the context
		// This is done so that the user can be accessed in the next middlewares
		c.Set("token", token)
		c.Set("roles", token.Claims.(jwt.MapClaims)["roles"])
		c.Set("user_id", token.Claims.(jwt.MapClaims)["sub"])

		// Next middleware on the way to hell (or the actual request)
		c.Next()
	}
}

// parseToken parses the token sent by the client
func parseToken(c *gin.Context, clientToken string, key *rsa.PublicKey) (bool, *jwt.Token) {
	// Parse the token sent by the client
	parsedToken, err := jwt.Parse(clientToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	// Check if there was an error while parsing the token
	if err != nil {
		if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Token Signature"})
			c.Abort()
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad Request"})
			c.Abort()
		}
		return false, nil
	}

	// Check if the token is valid
	if !parsedToken.Valid {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid Token"})
		c.Abort()
		return false, nil
	}

	// return true if the token is valid
	return true, parsedToken
}

// extractTokenFromContext extracts the token from the Authorization header
func extractTokenFromContext(c *gin.Context) string {
	// Get the Authorization header
	clientToken := c.GetHeader("Authorization")
	if clientToken == "" {
		log.Error().Msg("Authorization token was not provided")
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Authorization Token is required"})
		c.Abort()
		return ""
	}

	// Extract the token from the Authorization header
	extractedToken := strings.Split(clientToken, "Bearer ")
	if len(extractedToken) == 2 {
		return strings.TrimSpace(extractedToken[1])
	}

	// Incorrect Format of Authorization Token
	log.Error().Msg("Incorrect Format of Authorization Token")
	c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Incorrect Format of Authorization Token"})
	c.Abort()
	return ""
}

// parsePublicKey parses the public key from the SecretKey
func parsePublicKey(secret string) (*rsa.PublicKey, error) {
	return jwt.ParseRSAPublicKeyFromPEM([]byte(secret))
}
