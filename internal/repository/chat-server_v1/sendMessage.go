package chatrepo

import (
	"context"
	"log"

	"github.com/markgenuine/chat-server/internal/client/db"
	"github.com/markgenuine/chat-server/internal/model"
)

func (r *repo) SendMessage(ctx context.Context, message *model.Message) error {
	query, args, err := r.sq.Insert(chatsMessages).
		Columns(chatsMessagesChatID, chatsMessagesUserID, chatsMessagesBody, chatsMessagesTime).
		Values(message.ChatID, message.From, message.Text, message.Timestamp).
		ToSql()
	if err != nil {
		log.Printf("failed to build query insert message: %s", err.Error())
		return err
	}

	q := db.Query{
		Name:     "chat_repository.SendMessage",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		log.Printf("failed insert message to db: %s", err.Error())
		return err
	}

	return nil
}
