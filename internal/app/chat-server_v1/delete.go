package chat_server_v1

import (
	"context"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/markgenuine/chat-server/pkg/chat_server_v1"
)

// Delete chat
func (s *ChatServer) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	fmt.Printf("Delete chat with ID: %d", req.GetId())

	query, args, err := s.sq.Delete(ChatsMessages).
		Where(sq.Eq{ChatsMessagesChatID: req.GetId()}).
		ToSql()
	if err != nil {
		log.Printf("failed to build query for delete messages of chatId: %s", err.Error())
		return nil, err
	}

	_, err = s.poolDB.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to delete messages of chatId: %s", err.Error())
		return nil, err
	}

	query, args, err = s.sq.Delete(ChatsUsers).
		Where(sq.Eq{ChatsUsersChatID: req.GetId()}).
		ToSql()
	if err != nil {
		log.Printf("failed to build query for delete pair chatId and userId of chatId: %s", err.Error())
		return nil, err
	}

	_, err = s.poolDB.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to delete pair chatId and userId of chatId: %s", err.Error())
		return nil, err
	}

	query, args, err = s.sq.Delete(Chats).
		Where(sq.Eq{ChatsID: req.GetId()}).
		ToSql()
	if err != nil {
		log.Printf("failed to build query for delete chatId of chatId: %s", err.Error())
		return nil, err
	}

	_, err = s.poolDB.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to delete chatId of chatId: %s", err.Error())
		return nil, err
	}

	return &empty.Empty{}, nil
}
