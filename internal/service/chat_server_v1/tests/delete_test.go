package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/jackc/pgx/v4"
	"github.com/markgenuine/chat-server/internal/client/db"
	dbMocks "github.com/markgenuine/chat-server/internal/client/db/mocks"
	"github.com/markgenuine/chat-server/internal/client/db/transaction"
	"github.com/markgenuine/chat-server/internal/repository"
	repoMock "github.com/markgenuine/chat-server/internal/repository/mocks"
	chatServerService "github.com/markgenuine/chat-server/internal/service/chat_server_v1"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	t.Parallel()

	type chatServerRepoMockFunc func(mc *minimock.Controller) repository.ChatServerRepository
	type transactorMockFunc func(mc *minimock.Controller) db.Transactor

	type args struct {
		ctx     context.Context
		request int64
	}

	var (
		ctx        = context.Background()
		mc         = minimock.NewController(t)
		serviceErr = fmt.Errorf("failed to delete chat")

		userID = int64(gofakeit.Uint64())
		opts   = pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		err                error
		chatServerRepoMock chatServerRepoMockFunc
		transactorMock     transactorMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx:     ctx,
				request: userID,
			},
			err: nil,
			chatServerRepoMock: func(mc *minimock.Controller) repository.ChatServerRepository {
				mock := repoMock.NewChatServerRepositoryMock(mc)
				mock.DeleteChatMessagesMock.Expect(minimock.AnyContext, userID).Return(nil)
				return mock
			},
			transactorMock: func(mc *minimock.Controller) db.Transactor {
				mock := dbMocks.NewTransactorMock(mc)
				txMock := dbMocks.NewTxMock(mc)
				mock.BeginTxMock.Expect(minimock.AnyContext, opts).Return(txMock, nil)
				txMock.CommitMock.Expect(minimock.AnyContext).Return(nil)
				return mock
			},
		},
		{
			name: "cancel",
			args: args{
				ctx:     ctx,
				request: userID,
			},
			err: serviceErr,
			chatServerRepoMock: func(mc *minimock.Controller) repository.ChatServerRepository {
				mock := repoMock.NewChatServerRepositoryMock(mc)
				mock.DeleteChatMessagesMock.Expect(minimock.AnyContext, userID).Return(serviceErr)
				return mock
			},
			transactorMock: func(mc *minimock.Controller) db.Transactor {
				mock := dbMocks.NewTransactorMock(mc)
				txMock := dbMocks.NewTxMock(mc)
				mock.BeginTxMock.Expect(minimock.AnyContext, opts).Return(txMock, nil)
				txMock.RollbackMock.Expect(minimock.AnyContext).Return(nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			chatServerRepoMock := tt.chatServerRepoMock(mc)
			txManagerMock := transaction.NewTransactionManager(tt.transactorMock(mc))
			service := chatServerService.NewService(chatServerRepoMock, txManagerMock)

			err := service.Delete(tt.args.ctx, tt.args.request)
			require.Equal(t, tt.err, err)
		})
	}

}
