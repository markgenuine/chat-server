package chat_server_v1

import (
	"context"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit/v6"
	desc "github.com/markgenuine/chat-server/pkg/chat_server_v1"
)

// Create new chat
func (s *ChatServer) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	fmt.Printf("Create chat with users: %s", req.GetUsernames())

	query, args, err := s.sq.Insert(Chats).
		Columns(ChatsID).
		Values(squirrel.Expr("DEFAULT")).
		Suffix(fmt.Sprintf("RETURNING %s", ChatsID)).
		ToSql()

	if err != nil {
		log.Printf("failed to build query create chat: %s", err.Error())
		return nil, err
	}

	var chatID int64
	err = s.poolDB.QueryRow(ctx, query, args...).Scan(&chatID)
	if err != nil {
		log.Printf("failed to insert chat: %s", err.Error())
		return nil, err
	}

	builderChatUser := s.sq.Insert(ChatsUsers).Columns(ChatsUsersChatID, ChatsUsersUserID)

	// TODO add get id users
	for range req.GetUsernames() {
		id := int64(gofakeit.Uint8())
		fmt.Println(id)
		builderChatUser = builderChatUser.Values(chatID, id)
	}

	query, args, err = builderChatUser.ToSql()
	if err != nil {
		log.Printf("failed to build query create pair chatID and userID: %s", err.Error())
		return nil, err
	}

	_, err = s.poolDB.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to create pair chatId and userId: %s", err.Error())
		return nil, err
	}

	return &desc.CreateResponse{
		Id: chatID,
	}, nil
}
