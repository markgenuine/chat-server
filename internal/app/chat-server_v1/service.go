package chat_server_v1

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	desc "github.com/markgenuine/chat-server/pkg/chat_server_v1"
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

// ChatServer - type for proto implementation
type ChatServer struct {
	desc.UnimplementedChatServerV1Server

	poolDB *pgx.Conn
	sq     squirrel.StatementBuilderType
}

// NewChatServer create proto interface implementation
func NewChatServer(conn *pgx.Conn, sqIn squirrel.StatementBuilderType) *ChatServer {
	return &ChatServer{poolDB: conn, sq: sqIn}
}
