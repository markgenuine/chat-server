package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/markgenuine/chat-server/internal/model"
	"github.com/markgenuine/chat-server/internal/repository"
	repoMock "github.com/markgenuine/chat-server/internal/repository/mocks"
	chatServerService "github.com/markgenuine/chat-server/internal/service/chat_server_v1"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	type chatServerRepoMockFunc func(mc *minimock.Controller) repository.ChatServerRepository

	type args struct {
		ctx     context.Context
		request *model.Chat
	}

	var (
		ctx        = context.Background()
		mc         = minimock.NewController(t)
		serviceErr = fmt.Errorf("failed to create chat")

		userID    = int64(gofakeit.Uint64())
		usernames = []string{gofakeit.BeerName(), gofakeit.BeerName(), gofakeit.BeerName()}
		request   = &model.Chat{Usernames: usernames}
		response  = userID
	)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		chatServerRepoMock chatServerRepoMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx:     ctx,
				request: request,
			},
			want: response,
			err:  nil,
			chatServerRepoMock: func(mc *minimock.Controller) repository.ChatServerRepository {
				mock := repoMock.NewChatServerRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, request).Return(response, nil)
				return mock
			},
		},
		{
			name: "cancel",
			args: args{
				ctx:     ctx,
				request: request,
			},
			want: 0,
			err:  serviceErr,
			chatServerRepoMock: func(mc *minimock.Controller) repository.ChatServerRepository {
				mock := repoMock.NewChatServerRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, request).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			chatServerRepoMock := tt.chatServerRepoMock(mc)
			service := chatServerService.NewService(chatServerRepoMock, nil)

			result, err := service.Create(tt.args.ctx, tt.args.request)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
