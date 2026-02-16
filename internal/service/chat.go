package service

import (
	"context"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgxpool"

	v1 "moonshine-server/gen/moonshine/v1"
	"moonshine-server/gen/moonshine/v1/moonshinev1connect"
)

// ChatService implements the ChatServiceHandler interface.
type ChatService struct {
	moonshinev1connect.UnimplementedChatServiceHandler
	db *pgxpool.Pool
}

// NewChatService creates a new ChatService.
func NewChatService(db *pgxpool.Pool) *ChatService {
	return &ChatService{db: db}
}

// SendMessage inserts a new message into the channel.
func (s *ChatService) SendMessage(
	ctx context.Context,
	req *connect.Request[v1.SendMessageRequest],
) (*connect.Response[v1.SendMessageResponse], error) {
	// TODO: Extract user ID from auth context
	// TODO: Insert message into database
	// TODO: Broadcast to stream subscribers
	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}

// GetMessages retrieves paginated message history for a channel.
func (s *ChatService) GetMessages(
	ctx context.Context,
	req *connect.Request[v1.GetMessagesRequest],
) (*connect.Response[v1.GetMessagesResponse], error) {
	// TODO: Query messages from database with cursor pagination
	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}

// StreamMessages opens a server-streaming connection for real-time messages.
func (s *ChatService) StreamMessages(
	ctx context.Context,
	req *connect.Request[v1.StreamMessagesRequest],
	stream *connect.ServerStream[v1.StreamMessagesResponse],
) error {
	// TODO: Subscribe to channel messages and stream to client
	// TODO: Use a pub/sub mechanism (channels, Redis, etc.)
	return connect.NewError(connect.CodeUnimplemented, nil)
}
