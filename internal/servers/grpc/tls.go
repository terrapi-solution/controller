package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc/credentials"
	"os"
)

// getTransportCredentials returns the transport credentials for the gRPC server
func (s *GrpcServer) getTransportCredentials(certFile, keyFile, caFile string) (credentials.TransportCredentials, error) {
	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load servers certification: %w", err)
	}

	data, err := os.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("faild to read CA certificate: %w", err)
	}

	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(data) {
		return nil, fmt.Errorf("unable to append the CA certificate to CA pool")
	}

	tlsConfig := &tls.Config{
		ClientAuth:       tls.RequireAndVerifyClientCert,
		Certificates:     []tls.Certificate{certificate},
		ClientCAs:        pool,
		CurvePreferences: s.curves(),
		CipherSuites:     s.ciphers(),
	}
	return credentials.NewTLS(tlsConfig), nil
}

// Curves provides optionally a list of secure curves.
func (s *GrpcServer) curves() []tls.CurveID {
	if s.config.StrictCurves {
		return []tls.CurveID{
			tls.CurveP521,
			tls.CurveP384,
			tls.CurveP256,
		}
	}

	return nil
}

// Ciphers provides optionally a list of secure ciphers.
func (s *GrpcServer) ciphers() []uint16 {
	if s.config.StrictCiphers {
		return []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		}
	}

	return nil
}
