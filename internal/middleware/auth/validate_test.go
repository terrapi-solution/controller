package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"math/big"
	"net/http"
	"net/http/httptest"
	"time"

	"testing"
)

func TestValidateToken_ValidToken(t *testing.T) {
	validCertificate, validPrivateKey := generateEphemeralCertificate()

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(ValidateToken(validCertificate))
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})

	tokenString := createValidateToken(validPrivateKey)
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestValidateToken_InvalidTokenSignature(t *testing.T) {
	validCertificate, _ := generateEphemeralCertificate()
	_, invalidPrivateKey := generateEphemeralCertificate()

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(ValidateToken(validCertificate))
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})

	tokenString := createValidateToken(invalidPrivateKey)
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, w.Code)
	}
}

func TestValidateToken_NoTokenProvided(t *testing.T) {
	validCertificate, _ := generateEphemeralCertificate()

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(ValidateToken(validCertificate))
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, w.Code)
	}
}

func TestValidateToken_InvalidTokenFormat(t *testing.T) {
	validCertificate, _ := generateEphemeralCertificate()

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(ValidateToken(validCertificate))
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "InvalidTokenFormat")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
}

// generateEphemeralCertificate generates a self-signed certificate and private key
func generateEphemeralCertificate() (string, *rsa.PrivateKey) {
	// Generate a new RSA private key
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	// Create a template for the certificate
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"TerrAPI"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	// Create the certificate
	certBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)

	// Encode the certificate to PEM format
	certPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})

	return string(certPEM), privateKey
}

// createValidateToken creates a JWT token for testing purposes
func createValidateToken(privateKey *rsa.PrivateKey) string {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{})
	tokenString, _ := token.SignedString(privateKey)
	return tokenString
}
