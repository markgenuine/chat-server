package chat_server_v1

import desc "github.com/markgenuine/chat-server/pkg/chat_server_v1"

// ChatServer - type for proto implementation
type ChatServer struct {
	desc.UnimplementedChatServerV1Server
}

// NewChatServer create proto interface implementation
func NewChatServer() *ChatServer {
	return &ChatServer{}
}
