package service

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/terrapi-solution/controller/internal/config"
	"math/big"
	"net/http"
	"strings"
	"time"
)

// Auth represents an authentication service.
type Auth struct {
	authority  string
	publicKeys map[string]*rsa.PublicKey
	lastUpdate *time.Time
}

// NewAuthService creates a new instance of the Auth service.
func NewAuthService(cfg *config.Config) (*Auth, error) {
	return &Auth{
		authority: cfg.Auth.Authority,
	}, nil
}

// GetPublicKeys retrieves the RSA public keys from the JWK set.
func (a *Auth) GetPublicKeys() map[string]*rsa.PublicKey {
	if a.lastUpdate == nil || time.Since(*a.lastUpdate) > 6*time.Hour {

		jsonWebTokenAddress := a.authority + "/protocol/openid-connect/certs"
		result, err := a.getJWKSet(jsonWebTokenAddress)
		if err != nil && a.lastUpdate == nil {
			panic(fmt.Errorf("unable to retrieve public keys: %v", err))
		} else if err != nil && a.lastUpdate != nil {
			fmt.Printf("unable to update public keys: %v", err)
		} else {
			a.publicKeys = result
			*a.lastUpdate = time.Now().UTC()
			*a.lastUpdate = time.Now()
		}
	}
	return a.publicKeys
}

// ValidateToken parses and validates a JWT token string using the RSA public keys.
func (a *Auth) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Extract the key ID (kid) from the token header
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("kid not found in token header")
		}

		// Retrieve the public key from the JWK set using the key ID
		publicKey, keyExists := a.publicKeys[kid]
		if !keyExists {
			return nil, fmt.Errorf("public key not found for kid: %v", kid)
		}

		return publicKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return token, nil
}

// GetJWKSet retrieves the JWK set from the specified URL and returns a map of RSA public keys.
func (a *Auth) getJWKSet(url string) (map[string]*rsa.PublicKey, error) {
	// Make the GET request
	response, err := http.Get(url)
	if err != nil || response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			fmt.Printf("error closing response body: %v", cerr)
		}
	}()

	// Decode the JSON response
	var jwkSet struct {
		Keys []struct {
			Kid string   `json:"kid"`
			N   string   `json:"n"`
			E   string   `json:"e"`
			X5C []string `json:"x5c"`
		} `json:"keys"`
	}
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&jwkSet); err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %v", err)
	}

	// Create a map to store RSA public keys
	jwkMap := make(map[string]*rsa.PublicKey)

	// Iterate through each key in the JWK set
	for _, key := range jwkSet.Keys {
		// Decode base64url-encoded modulus (N) and exponent (E)
		modulus, err := a.decodeBase64URL(key.N)
		if err != nil {
			return nil, fmt.Errorf("error decoding modulus: %v", err)
		}

		exponent, err := a.decodeBase64URL(key.E)
		if err != nil {
			return nil, fmt.Errorf("error decoding exponent: %v", err)
		}

		// Create RSA public key
		pubKey := &rsa.PublicKey{
			N: modulus,
			E: int(exponent.Int64()),
		}

		// Store the public key in the map using the key ID (Kid)
		jwkMap[key.Kid] = pubKey
	}

	return jwkMap, nil
}

// decodeBase64URL decodes a base64url-encoded string and returns a big.Int
func (a *Auth) decodeBase64URL(input string) (*big.Int, error) {
	// Convert base64url to base64
	base64Str := strings.ReplaceAll(input, "-", "+")
	base64Str = strings.ReplaceAll(base64Str, "_", "/")

	// Pad the base64 string with "="
	switch len(base64Str) % 4 {
	case 2:
		base64Str += "=="
	case 3:
		base64Str += "="
	}

	// Decode base64 string
	data, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, fmt.Errorf("error decoding base64 string: %v", err)
	}

	// Convert bytes to big.Int
	result := new(big.Int).SetBytes(data)
	return result, nil
}
