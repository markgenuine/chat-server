package chatservice

import (
	"github.com/markgenuine/chat-server/internal/repository"
	def "github.com/markgenuine/chat-server/internal/service"
	"github.com/markgenuine/platform_common/pkg/db"
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
