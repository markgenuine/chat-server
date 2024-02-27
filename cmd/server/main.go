package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v5"
	chatServerV1 "github.com/markgenuine/chat-server/internal/app/chat-server_v1"
	"github.com/markgenuine/chat-server/internal/config"
	"github.com/markgenuine/chat-server/internal/config/env"
	desc "github.com/markgenuine/chat-server/pkg/chat_server_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	fmt.Println(configPath)
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %s ", err.Error())
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %s ", err.Error())
	}

	list, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgx.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err.Error())
	}

	defer func() {
		err = pool.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatServerV1Server(
		s,
		chatServerV1.NewChatServer(pool, sq.StatementBuilder.PlaceholderFormat(sq.Dollar)),
	)

	log.Printf("server listerning at %v", list.Addr())

	if err = s.Serve(list); err != nil {
		log.Fatalf("failed to serv %v", err)
	}
}
