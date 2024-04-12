package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/jackc/pgx/v4"
	"github.com/markgenuine/chat-server/internal/repository"
	repoMock "github.com/markgenuine/chat-server/internal/repository/mocks"
	chatServerService "github.com/markgenuine/chat-server/internal/service/chat_server_v1"
	"github.com/markgenuine/platform_common/pkg/db"
	"github.com/markgenuine/platform_common/pkg/db/mocks"
	"github.com/markgenuine/platform_common/pkg/db/pg"
	"github.com/markgenuine/platform_common/pkg/db/transaction"
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
		userID     = int64(gofakeit.Uint64())
		txOpts     = pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	)

	txMock := mocks.NewTxMock(mc)
	ctxWithTx := pg.MakeContextTx(ctx, txMock)

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
				mock.DeleteChatMessagesMock.Expect(ctxWithTx, userID).Return(nil)
				mock.DeleteChatsUsersMock.Expect(ctxWithTx, userID).Return(nil)
				mock.DeleteChatUsersMock.Expect(ctxWithTx, userID).Return(nil)
				return mock
			},
			transactorMock: func(mc *minimock.Controller) db.Transactor {
				mock := mocks.NewTransactorMock(mc)
				mock.BeginTxMock.Expect(ctx, txOpts).Return(txMock, nil)
				txMock.CommitMock.Expect(ctxWithTx).Return(nil)

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
				mock.DeleteChatMessagesMock.Expect(ctxWithTx, userID).Return(serviceErr)
				return mock
			},
			transactorMock: func(mc *minimock.Controller) db.Transactor {
				mock := mocks.NewTransactorMock(mc)
				mock.BeginTxMock.Expect(ctx, txOpts).Return(txMock, nil)
				txMock.RollbackMock.Expect(ctxWithTx).Return(nil)

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

			err := service.Delete(ctx, tt.args.request)
			require.Equal(t, tt.err, err)
		})
	}
}
