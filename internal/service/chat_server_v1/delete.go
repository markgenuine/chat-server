package chat_server_v1

import (
	"context"
	"errors"
	"log"
)

func (s *service) Delete(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTx := s.chatRepository.Delete(ctx, id)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		log.Print(err)
		return errors.New("failed to delete chat")
	}
	return nil
}
