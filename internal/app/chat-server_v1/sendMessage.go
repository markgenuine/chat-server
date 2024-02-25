package chat_server_v1

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/markgenuine/chat-server/pkg/chat_server_v1"
)

// SendMessage in chat
func (s *ChatServer) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	fmt.Printf("Message from: %s", req.GetFrom())
	fmt.Printf("Message text: %s", req.GetText())
	fmt.Printf("Message date: %s", req.GetTimestamp())

	query, args, err := s.sq.Insert(ChatsMessages).
		Columns(ChatsMessagesChatID, ChatsMessagesUserID, ChatsMessagesBody, ChatsMessagesTime).
		Values(req.GetChatId(), req.GetFrom(), req.GetText(), req.GetTimestamp()).
		ToSql()
	if err != nil {
		log.Printf("failed to build query insert message: %s", err.Error())
		return nil, err
	}

	_, err = s.poolDB.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed insert message to db: %s", err.Error())
		return nil, err
	}

	return &empty.Empty{}, nil
}
