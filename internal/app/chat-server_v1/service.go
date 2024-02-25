package chat_server_v1

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	desc "github.com/markgenuine/chat-server/pkg/chat_server_v1"
)

const (
	// Chats ...
	Chats = "chats"

	// ChatsID ...
	ChatsID = "id"

	// ChatsUsers ...
	ChatsUsers = "chats_users"

	// ChatsUsersChatID ..
	ChatsUsersChatID = "chat_id"

	// ChatsUsersUserID ...
	ChatsUsersUserID = "user_id"

	// ChatsMessages ...
	ChatsMessages = "chats_messages"

	// ChatsMessagesID ...
	ChatsMessagesID = "id"

	// ChatsMessagesChatID ...
	ChatsMessagesChatID = "chat_id"

	// ChatsMessagesUserID ...
	ChatsMessagesUserID = "user_id"

	// ChatsMessagesBody ...
	ChatsMessagesBody = "body"

	// ChatsMessagesTime ...
	ChatsMessagesTime = "time"
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
