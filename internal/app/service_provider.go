package app

import (
	"context"
	"log"

	api "github.com/markgenuine/chat-server/internal/api/chat-server"
	"github.com/markgenuine/chat-server/internal/config"
	"github.com/markgenuine/chat-server/internal/config/env"
	"github.com/markgenuine/chat-server/internal/repository"
	chatServerRepo "github.com/markgenuine/chat-server/internal/repository/chat-server_v1"
	"github.com/markgenuine/chat-server/internal/service"
	chatServerService "github.com/markgenuine/chat-server/internal/service/chat_server_v1"
	"github.com/markgenuine/platform_common/pkg/closer"
	"github.com/markgenuine/platform_common/pkg/db"
	"github.com/markgenuine/platform_common/pkg/db/pg"
	"github.com/markgenuine/platform_common/pkg/db/transaction"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient  db.Client
	txManager db.TxManager

	chatServerRepository repository.ChatServerRepository
	chatService          service.ChatService
	chatServerImpl       *api.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// PGConfig ...
func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}
		s.pgConfig = cfg
	}

	return s.pgConfig
}

// GRPCConfig ...
func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}
		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %s", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err)
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.ChatServerRepository {
	if s.chatServerRepository == nil {
		s.chatServerRepository = chatServerRepo.NewRepository(s.DBClient(ctx))
	}

	return s.chatServerRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatServerService.NewService(s.UserRepository(ctx), s.TxManager(ctx))
	}

	return s.chatService
}

func (s *serviceProvider) ChatServerImpl(ctx context.Context) *api.Implementation {
	if s.chatServerImpl == nil {
		s.chatServerImpl = api.NewImplementation(s.ChatService(ctx))
	}

	return s.chatServerImpl
}
