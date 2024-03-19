package chatserver

import (
	"github.com/markgenuine/chat-server/internal/service"

	desc "github.com/markgenuine/chat-server/pkg/chat_server_v1"
)

// Implementation - type for proto implementation
type Implementation struct {
	desc.UnimplementedChatServerV1Server
	chatServer service.ChatService
}

// NewImplementation create proto interface implementation
func NewImplementation(chatServer service.ChatService) *Implementation {
	return &Implementation{
		chatServer: chatServer,
	}
}
