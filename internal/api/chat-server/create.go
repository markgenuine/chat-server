package chat_server

import (
	"context"
	"log"

	"github.com/markgenuine/chat-server/internal/converter"
	desc "github.com/markgenuine/chat-server/pkg/chat_server_v1"
)

// Create new chat
func (s *Implementation) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := s.chatServer.Create(ctx, converter.CreateToServiceFromChat(request))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted chat with id: %d", id)

	return converter.CreateToChatFromService(id), nil
}
