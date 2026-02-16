package service

import (
	"context"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgxpool"

	v1 "moonshine-server/gen/moonshine/v1"
	"moonshine-server/gen/moonshine/v1/moonshinev1connect"
)

// UserService implements the UserServiceHandler interface.
type UserService struct {
	moonshinev1connect.UnimplementedUserServiceHandler
	db *pgxpool.Pool
}

// NewUserService creates a new UserService.
func NewUserService(db *pgxpool.Pool) *UserService {
	return &UserService{db: db}
}

// GetUser retrieves a user by ID.
func (s *UserService) GetUser(
	ctx context.Context,
	req *connect.Request[v1.GetUserRequest],
) (*connect.Response[v1.GetUserResponse], error) {
	// TODO: Query user from database
	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}

// UpdateUser updates the authenticated user's profile.
func (s *UserService) UpdateUser(
	ctx context.Context,
	req *connect.Request[v1.UpdateUserRequest],
) (*connect.Response[v1.UpdateUserResponse], error) {
	// TODO: Extract user ID from auth context
	// TODO: Update user in database
	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}

// ListUsers returns all users.
func (s *UserService) ListUsers(
	ctx context.Context,
	req *connect.Request[v1.ListUsersRequest],
) (*connect.Response[v1.ListUsersResponse], error) {
	// TODO: Query all users from database
	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}
