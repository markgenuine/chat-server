package chat_server_v1

import (
	"github.com/markgenuine/chat-server/internal/client/db"
	"github.com/markgenuine/chat-server/internal/repository"
	def "github.com/markgenuine/chat-server/internal/service"
)

var _ def.ChatService = (*service)(nil)

type service struct {
	chatRepository repository.ChatServerRepository
	txManager      db.TxManager
}

// NewService ...
func NewService(chatRepository repository.ChatServerRepository, txManager db.TxManager) *service {
	return &service{
		chatRepository: chatRepository,
		txManager:      txManager,
	}
}
