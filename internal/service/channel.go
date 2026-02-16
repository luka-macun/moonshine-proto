package service

import (
	"context"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgxpool"

	v1 "moonshine-server/gen/moonshine/v1"
	"moonshine-server/gen/moonshine/v1/moonshinev1connect"
)

// ChannelService implements the ChannelServiceHandler interface.
type ChannelService struct {
	moonshinev1connect.UnimplementedChannelServiceHandler
	db *pgxpool.Pool
}

// NewChannelService creates a new ChannelService.
func NewChannelService(db *pgxpool.Pool) *ChannelService {
	return &ChannelService{db: db}
}

// CreateChannel creates a new chat channel.
func (s *ChannelService) CreateChannel(
	ctx context.Context,
	req *connect.Request[v1.CreateChannelRequest],
) (*connect.Response[v1.CreateChannelResponse], error) {
	// TODO: Extract user ID from auth context
	// TODO: Insert channel into database
	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}

// GetChannel retrieves a single channel by ID.
func (s *ChannelService) GetChannel(
	ctx context.Context,
	req *connect.Request[v1.GetChannelRequest],
) (*connect.Response[v1.GetChannelResponse], error) {
	// TODO: Query channel from database
	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}

// ListChannels returns all channels.
func (s *ChannelService) ListChannels(
	ctx context.Context,
	req *connect.Request[v1.ListChannelsRequest],
) (*connect.Response[v1.ListChannelsResponse], error) {
	// TODO: Query all channels from database
	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}

// UpdateChannel updates a channel's name and description.
func (s *ChannelService) UpdateChannel(
	ctx context.Context,
	req *connect.Request[v1.UpdateChannelRequest],
) (*connect.Response[v1.UpdateChannelResponse], error) {
	// TODO: Update channel in database
	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}

// DeleteChannel removes a channel.
func (s *ChannelService) DeleteChannel(
	ctx context.Context,
	req *connect.Request[v1.DeleteChannelRequest],
) (*connect.Response[v1.DeleteChannelResponse], error) {
	// TODO: Delete channel from database
	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}
