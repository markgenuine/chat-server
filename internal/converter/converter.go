package converter

import (
	"strconv"

	"github.com/markgenuine/chat-server/internal/model"
	desc "github.com/markgenuine/chat-server/pkg/chat_server_v1"
)

// CreateToServiceFromChat ...
func CreateToServiceFromChat(request *desc.CreateRequest) *model.Chat {
	return &model.Chat{
		Usernames: request.GetUsernames(),
	}
}

// CreateToChatFromService ...
func CreateToChatFromService(id int64) *desc.CreateResponse {
	return &desc.CreateResponse{
		Id: id,
	}
}

// DeleteToServiceFromChat ...
func DeleteToServiceFromChat(request *desc.DeleteRequest) int64 {
	return request.GetId()
}

// SendMessageToServiceFromChat ...
func SendMessageToServiceFromChat(request *desc.SendMessageRequest) *model.Message {
	userID, _ := strconv.Atoi(request.From)
	return &model.Message{
		ChatID:    request.ChatId,
		From:      int64(userID),
		Text:      request.From,
		Timestamp: request.Timestamp.AsTime(),
	}
}
