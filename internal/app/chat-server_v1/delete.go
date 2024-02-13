package chat_server_v1

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/markgenuine/chat-server/pkg/chat_server_v1"
)

// Delete chat
func (s *ChatServer) Delete(ctx context.Context, request *desc.DeleteRequest) (*empty.Empty, error) {
	_ = ctx
	fmt.Printf("Delete chat with ID: %d", request.GetId())

	return &empty.Empty{}, nil
}
