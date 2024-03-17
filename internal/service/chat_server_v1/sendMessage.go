package chat_server_v1

import (
	"context"
	"errors"
	"log"

	"github.com/markgenuine/chat-server/internal/model"
)

func (s *service) SendMessage(ctx context.Context, message *model.Message) error {

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTx := s.chatRepository.SendMessage(ctx, message)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		log.Print(err)
		return errors.New("failed to send message")
	}
	return nil
}
