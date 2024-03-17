package chat_server_v1

import (
	"context"
	"errors"
	"log"

	"github.com/markgenuine/chat-server/internal/model"
)

func (s *service) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	var id int64

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.chatRepository.Create(ctx, chat)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		log.Print(err)
		return 0, errors.New("failed to create chat")
	}

	return id, nil
}
