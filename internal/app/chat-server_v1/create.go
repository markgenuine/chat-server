package chat_server_v1

import (
	"context"
	"fmt"

	desc "github.com/markgenuine/chat-server/pkg/chat_server_v1"
)

// Create new chat
func (s *ChatServer) Create(_ context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	fmt.Printf("Create chat with users: %s", request.GetUsernames())

	return &desc.CreateResponse{
		Id: 0,
	}, nil
}
