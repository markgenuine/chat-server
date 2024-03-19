package chatservice

import (
	"context"
	"errors"
	"log"

	"github.com/markgenuine/chat-server/internal/model"
)

func (s *service) SendMessage(ctx context.Context, message *model.Message) error {
	err := s.chatRepository.SendMessage(ctx, message)
	if err != nil {
		log.Print(err)
		return errors.New("failed to send message")
	}

	return nil
}
