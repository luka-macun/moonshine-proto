package service

import (
	"context"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgxpool"

	v1 "moonshine-server/gen/moonshine/v1"
	"moonshine-server/gen/moonshine/v1/moonshinev1connect"
)

// AuthService implements the AuthServiceHandler interface.
type AuthService struct {
	moonshinev1connect.UnimplementedAuthServiceHandler
	db *pgxpool.Pool
}

// NewAuthService creates a new AuthService.
func NewAuthService(db *pgxpool.Pool) *AuthService {
	return &AuthService{db: db}
}

// Register creates a new user account.
func (s *AuthService) Register(
	ctx context.Context,
	req *connect.Request[v1.RegisterRequest],
) (*connect.Response[v1.RegisterResponse], error) {
	// TODO: Hash password with bcrypt
	// TODO: Insert user into database
	// TODO: Generate JWT token
	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}

// Login authenticates a user and returns a token.
func (s *AuthService) Login(
	ctx context.Context,
	req *connect.Request[v1.LoginRequest],
) (*connect.Response[v1.LoginResponse], error) {
	// TODO: Look up user by email
	// TODO: Compare password hash
	// TODO: Generate JWT token
	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}
