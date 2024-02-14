package chat_server_v1

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/markgenuine/chat-server/pkg/chat_server_v1"
)

// SendMessage in chat
func (s *ChatServer) SendMessage(_ context.Context, request *desc.SendMessageRequest) (*empty.Empty, error) {
	fmt.Printf("Message from: %s", request.GetFrom())
	fmt.Printf("Message text: %s", request.GetText())
	fmt.Printf("Message date: %s", request.GetTimestamp())

	return &empty.Empty{}, nil
}
