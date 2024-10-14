package service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type (
	// Validator defines an interface for token validation.
	// This is satisfied by auth service.
	Validator interface {
		ValidateToken(tokenString string) (*jwt.Token, error)
	}

	AuthInterceptor struct {
		validator Validator
	}
)

// NewAuthInterceptorService creates a new instance of the GRPC interceptor service.
func NewAuthInterceptorService(validator Validator) (*AuthInterceptor, error) {
	if validator == nil {
		return nil, errors.New("validator cannot be nil")
	}
	return &AuthInterceptor{validator: validator}, nil
}

func (i *AuthInterceptor) UnaryAuthMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// get metadata object
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
	}

	// extract token from authorization header
	token := md["authorization"]
	if len(token) == 0 {
		return nil, status.Error(codes.Unauthenticated, "authorization token is not provided")
	}

	// validate token
	_, err := i.validator.ValidateToken(token[0])
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}

	// call our handler
	return handler(ctx, req)
}
