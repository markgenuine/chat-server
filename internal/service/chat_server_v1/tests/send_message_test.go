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

func TestSendMessage(t *testing.T) {
	t.Parallel()

	type chatServerRepoMockFunc func(mc *minimock.Controller) repository.ChatServerRepository

	type args struct {
		ctx     context.Context
		request *model.Message
	}

	var (
		ctx        = context.Background()
		mc         = minimock.NewController(t)
		serviceErr = fmt.Errorf("failed to send message")

		chatID     = int64(gofakeit.Uint64())
		fromUserID = int64(gofakeit.Uint64())
		text       = gofakeit.BeerAlcohol()
		timestamp  = gofakeit.Date()
		request    = &model.Message{
			ChatID:    chatID,
			From:      fromUserID,
			Text:      text,
			Timestamp: timestamp,
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		err                error
		chatServerRepoMock chatServerRepoMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx:     ctx,
				request: request,
			},
			err: nil,
			chatServerRepoMock: func(mc *minimock.Controller) repository.ChatServerRepository {
				mock := repoMock.NewChatServerRepositoryMock(mc)
				mock.SendMessageMock.Expect(ctx, request).Return(nil)
				return mock
			},
		},
		{
			name: "cancel",
			args: args{
				ctx:     ctx,
				request: request,
			},
			err: serviceErr,
			chatServerRepoMock: func(mc *minimock.Controller) repository.ChatServerRepository {
				mock := repoMock.NewChatServerRepositoryMock(mc)
				mock.SendMessageMock.Expect(ctx, request).Return(serviceErr)
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

			err := service.SendMessage(tt.args.ctx, tt.args.request)
			require.Equal(t, tt.err, err)
		})
	}
}
