package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	chatserver "github.com/markgenuine/chat-server/internal/api/chat-server"
	"github.com/markgenuine/chat-server/internal/converter"
	"github.com/markgenuine/chat-server/internal/service"
	serviceMock "github.com/markgenuine/chat-server/internal/service/mocks"
	desc "github.com/markgenuine/chat-server/pkg/chat_server_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()

	type chatServerServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx     context.Context
		request *desc.SendMessageRequest
	}

	var (
		ctx        = context.Background()
		mc         = minimock.NewController(t)
		serviceErr = fmt.Errorf("service error")

		chatID    = int64(gofakeit.Uint64())
		from      = gofakeit.Name()
		text      = gofakeit.BeerAlcohol()
		timestamp = gofakeit.Date()
		request   = &desc.SendMessageRequest{
			ChatId:    chatID,
			From:      from,
			Text:      text,
			Timestamp: timestamppb.New(timestamp),
		}

		inputData = converter.SendMessageToServiceFromChat(request)
		response  = &emptypb.Empty{}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name                  string
		args                  args
		err                   error
		chatServerServiceMock chatServerServiceMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx:     ctx,
				request: request,
			},
			err: nil,
			chatServerServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMock.NewChatServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, inputData).Return(nil)
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
			chatServerServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMock.NewChatServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, inputData).Return(serviceErr)
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

			result, err := api.SendMessage(tt.args.ctx, tt.args.request)
			require.Equal(t, tt.err, err)
			require.Equal(t, response, result)
		})
	}
}
