package chatserver

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/markgenuine/chat-server/internal/converter"
	desc "github.com/markgenuine/chat-server/pkg/chat_server_v1"
)

// Delete ...
func (s *Implementation) Delete(ctx context.Context, request *desc.DeleteRequest) (*empty.Empty, error) {
	err := s.chatServer.Delete(ctx, converter.DeleteToServiceFromChat(request))

	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, err

}
