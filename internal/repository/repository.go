package repository

import (
	"context"

	"github.com/markgenuine/chat-server/internal/model"
)

// ChatServerRepository ...
type ChatServerRepository interface {
	Create(ctx context.Context, chat *model.Chat) (int64, error)
	DeleteChatMessages(ctx context.Context, id int64) error
	DeleteChatsUsers(ctx context.Context, id int64) error
	DeleteChatUsers(ctx context.Context, id int64) error
	SendMessage(ctx context.Context, message *model.Message) error
}
