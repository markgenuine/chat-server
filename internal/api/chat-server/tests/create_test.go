package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/markgenuine/chat-server/internal/api/chat-server"
	"github.com/markgenuine/chat-server/internal/converter"
	"github.com/markgenuine/chat-server/internal/service"
	serviceMock "github.com/markgenuine/chat-server/internal/service/mocks"
	desc "github.com/markgenuine/chat-server/pkg/chat_server_v1"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	type chatServerServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx     context.Context
		request *desc.CreateRequest
	}

	var (
		ctx        = context.Background()
		mc         = minimock.NewController(t)
		serviceErr = fmt.Errorf("service error")

		userID    = int64(gofakeit.Uint64())
		usernames = []string{gofakeit.BeerName(), gofakeit.BeerName(), gofakeit.BeerName()}
		inputData = &desc.CreateRequest{
			Usernames: usernames,
		}

		request  = converter.CreateToServiceFromChat(inputData)
		response = &desc.CreateResponse{Id: userID}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name                  string
		args                  args
		want                  *desc.CreateResponse
		err                   error
		chatServerServiceMock chatServerServiceMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx:     ctx,
				request: inputData,
			},
			want: response,
			err:  nil,
			chatServerServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMock.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx, request).Return(userID, nil)
				return mock
			},
		},
		{
			name: "cancel",
			args: args{
				ctx:     ctx,
				request: inputData,
			},
			want: nil,
			err:  serviceErr,
			chatServerServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMock.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx, request).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			chatServerServiceMock := tt.chatServerServiceMock(mc)
			api := chatserver.NewImplementation(chatServerServiceMock)

			result, err := api.Create(tt.args.ctx, tt.args.request)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
