package chatrepo

import (
	"context"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/markgenuine/chat-server/internal/model"
	"github.com/markgenuine/platform_common/pkg/db"
)

func (r *repo) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	query, args, err := r.sq.Insert(chats).
		Columns(chatsID).
		Values(squirrel.Expr("DEFAULT")).
		Suffix(fmt.Sprintf("RETURNING %s", chatsID)).
		ToSql()

	if err != nil {
		log.Printf("failed to build query create chat: %s", err.Error())
		return 0, err
	}

	q := db.Query{
		Name:     "chat_repository.Create_chats",
		QueryRaw: query,
	}
	var chatID int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chatID)
	if err != nil {
		log.Printf("failed to insert chat: %s", err.Error())
		return 0, err
	}

	builderChatUser := r.sq.Insert(chatsUsers).Columns(chatsUsersChatID, chatsUsersUserID)
	for range chat.Usernames {
		newUUID, _ := uuid.NewUUID()
		builderChatUser = builderChatUser.Values(chatID, int64(newUUID.ID()))
	}

	query, args, err = builderChatUser.ToSql()
	if err != nil {
		log.Printf("failed to build query create pair chatID and userID: %s", err.Error())
		return 0, err
	}

	q = db.Query{
		Name:     "chat_repository.Create_chatUsers",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		log.Printf("failed to create pair chatId and userId: %s", err.Error())
		return 0, err
	}

	return chatID, nil
}
