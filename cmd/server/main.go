package main

import (
	"log"
	"net"

	chat_server_v1 "github.com/markgenuine/chat-server/internal/app/chat-server_v1"
	desc "github.com/markgenuine/chat-server/pkg/chat_server_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	hostGRPC = "localhost:50552"
)

func main() {
	list, err := net.Listen("tcp", hostGRPC)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatServerV1Server(s, chat_server_v1.NewChatServer())

	log.Printf("server listerning at %v", list.Addr())

	if err = s.Serve(list); err != nil {
		log.Fatalf("failed to serv %v", err)
	}
}
