package chatrepo

import (
	"github.com/Masterminds/squirrel"
	"github.com/markgenuine/chat-server/internal/repository"
	"github.com/markgenuine/platform_common/pkg/db"
)

const (
	chats               = "chats"
	chatsID             = "id"
	chatsUsers          = "chats_users"
	chatsUsersChatID    = "chat_id"
	chatsUsersUserID    = "user_id"
	chatsMessages       = "chats_messages"
	chatsMessagesChatID = "chat_id"
	chatsMessagesUserID = "user_id"
	chatsMessagesBody   = "body"
	chatsMessagesTime   = "time"
)

type repo struct {
	db db.Client
	sq squirrel.StatementBuilderType
}

// NewRepository ...
func NewRepository(db db.Client) repository.ChatServerRepository {
	return &repo{
		db: db,
		sq: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
