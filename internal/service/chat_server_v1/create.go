package chatservice

import (
	"context"
	"errors"
	"log"

	"github.com/markgenuine/chat-server/internal/model"
)

func (s *service) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	id, err := s.chatRepository.Create(ctx, chat)
	if err != nil {
		log.Print(err)
		return 0, errors.New("failed to create chat")
	}

	return id, nil
}
