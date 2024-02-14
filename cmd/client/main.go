package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/brianvoe/gofakeit"
	desc "github.com/markgenuine/chat-server/pkg/chat_server_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	hostGRPC = "localhost:50552"
)

func main() {
	conn, err := grpc.Dial(hostGRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed connect to server: %v", err)
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			log.Fatalf("failed close connect: %v", err)
		}
	}()

	c := desc.NewChatServerV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var userNames []string

	userNames = append(userNames, gofakeit.Name())
	userNames = append(userNames, gofakeit.Name())
	userNames = append(userNames, gofakeit.Name())

	newChat, err := c.Create(ctx, &desc.CreateRequest{Usernames: userNames})
	if err != nil {
		log.Printf("failed create chat: %v", err)
	}

	fmt.Printf("New chat ID: %d", newChat.GetId())
}
