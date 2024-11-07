package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/user"
	"net/http"
	"strings"
)

// userKey is the key used to store the user in the context
const userKey = "user"

// Handler checks if the JWT sent is valid or not.
func Handler(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the user JWT from the Authorization header
		userJWT := extractUserJWT(c)

		// Parse the public key from the SecretKey
		key, err := parsePublicKey(secret)
		if err != nil {
			log.Error().Err(err).Msg("Error parsing token")
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				handleError(c, http.StatusUnauthorized, "Invalid Token Signature")
			} else {
				handleError(c, http.StatusBadRequest, "Bad Request")
			}
		}

		// Validate the token sent by the client
		status, token := parseToken(c, userJWT, key)
		if !status {
			return
		}

		// Store the user in the context
		// This is done so that the user can be accessed in the next middlewares
		storeUserInContext(c, token)

		// Next middleware on the way to hell (or the actual request)
		c.Next()
	}
}

// extractToken extracts the token from the Authorization header
func extractUserJWT(c *gin.Context) string {
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
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			handleError(c, http.StatusUnauthorized, "Invalid Token Signature")
		} else {
			handleError(c, http.StatusBadRequest, "Bad Request")
		}
		return false, nil
	}

	// Check if the token is valid
	if !parsedToken.Valid {
		handleError(c, http.StatusUnauthorized, "Invalid Token")
		return false, nil
	}

	// return true if the token is valid
	return true, parsedToken
}

// handleError handles the error and sends the response to the client
func handleError(c *gin.Context, status int, message string) {
	log.Error().Msg(message)
	c.JSON(status, gin.H{"status": status, "message": message})
	c.Abort()
}

// storeUserInContext stores the user in the context
func storeUserInContext(c *gin.Context, token *jwt.Token) {
	// Get the claims from the token
	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Error().Msg("Error asserting token claims")
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Error asserting token claims"})
		c.Abort()
		return
	}

	// Get the subject from the token
	sub, err := mapClaims.GetSubject()
	if err != nil {
		log.Error().Msg("Error getting subject from token: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Error getting subject from token"})
		c.Abort()
		return
	}

	// Get the role from the token
	role, ok := mapClaims["role"].(string)
	if !ok {
		log.Error().Msg("Error getting role from token")
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Error getting role from token"})
		c.Abort()
		return
	}

	// Get the userName from the token
	userName, ok := mapClaims["preferred_username"].(string)
	if !ok {
		log.Error().Msg("Error getting userName from token")
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Error getting userName from token"})
		c.Abort()
		return
	}

	// Store the user in the context
	c.Set(userKey, &user.User{
		Id:       sub,
		UserName: userName,
		Role:     user.Role(role),
	})
}
