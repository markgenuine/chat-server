package chatrepo

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/markgenuine/platform_common/pkg/db"
)

func (r *repo) DeleteChatMessages(ctx context.Context, id int64) error {
	query, args, err := r.sq.Delete(chatsMessages).
		Where(sq.Eq{chatsMessagesChatID: id}).
		ToSql()
	if err != nil {
		log.Printf("failed to build query for delete messages of chatId: %s", err.Error())
		return err
	}

	q := db.Query{
		Name:     "chat_repository.Delete_chatMessage",
		QueryRaw: query,
	}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		log.Printf("failed to delete messages of chatId: %s", err.Error())
		return err
	}

	return err
}

func (r *repo) DeleteChatsUsers(ctx context.Context, id int64) error {
	query, args, err := r.sq.Delete(chatsUsers).
		Where(sq.Eq{chatsUsersChatID: id}).
		ToSql()
	if err != nil {
		log.Printf("failed to build query for delete pair chatId and userId of chatId: %s", err.Error())
		return err
	}

	q := db.Query{
		Name:     "chat_repository.Delete_chatsUsers",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		log.Printf("failed to delete pair chatId and userId of chatId: %s", err.Error())
		return err
	}

	return nil
}

func (r *repo) DeleteChatUsers(ctx context.Context, id int64) error {
	query, args, err := r.sq.Delete(chats).
		Where(sq.Eq{chatsID: id}).
		ToSql()
	if err != nil {
		log.Printf("failed to build query for delete chatId of chatId: %s", err.Error())
		return err
	}

	q := db.Query{
		Name:     "chat_repository.Delete_chats",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		log.Printf("failed to delete chatId of chatId: %s", err.Error())
		return err
	}

	return nil
}
