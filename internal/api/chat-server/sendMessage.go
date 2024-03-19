package chatserver

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/markgenuine/chat-server/internal/converter"
	desc "github.com/markgenuine/chat-server/pkg/chat_server_v1"
)

// SendMessage in chat
func (s *Implementation) SendMessage(ctx context.Context, request *desc.SendMessageRequest) (*empty.Empty, error) {
	err := s.chatServer.SendMessage(ctx, converter.SendMessageToServiceFromChat(request))
	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}
