package app

import (
	"context"
	"github.com/Slintox/chat_server/config"
	chatRepo "github.com/Slintox/chat_server/internal/repository/chat"
	chatService "github.com/Slintox/chat_server/internal/service/chat"
	"github.com/Slintox/chat_server/pkg/common/closer"
	"github.com/Slintox/chat_server/pkg/database/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type ServiceProvider interface {
	GetConfig() *config.Config
}

type serviceProvider struct {
	configPath string

	config   *config.Config
	pgClient postgres.Client

	chatRepo    chatRepo.Repository
	chatService chatService.Service
}

func newServiceProvider(configPath string) *serviceProvider {
	return &serviceProvider{
		configPath: configPath,
	}
}

func (s *serviceProvider) GetConfig() *config.Config {
	if s.config != nil {
		return s.config
	}

	var err error
	s.config, err = config.InitConfig(s.configPath)
	if err != nil {
		log.Fatalf("failed to init config: %s", err.Error())
		return nil
	}

	return s.config
}

func (s *serviceProvider) GetPostgresClient(ctx context.Context) postgres.Client {
	if s.pgClient != nil {
		return s.pgClient
	}

	pgCfg, err := pgxpool.ParseConfig(s.GetConfig().GetPostgresConfig().DSN)
	if err != nil {
		log.Fatalf("failed to get db config: %s", err.Error())
	}

	client, err := postgres.NewClient(ctx, pgCfg)
	if err != nil {
		log.Fatalf("failed to get postgres client: %s", err.Error())
	}

	err = client.Postgres().Ping(ctx)
	if err != nil {
		log.Fatalf("ping error: %s", err.Error())
	}
	closer.Add(client.Close)

	s.pgClient = client
	return s.pgClient
}

func (s *serviceProvider) GetChatRepository(ctx context.Context) chatRepo.Repository {
	if s.chatRepo != nil {
		return s.chatRepo
	}

	s.chatRepo = chatRepo.NewRepository(s.GetPostgresClient(ctx))
	return s.chatRepo
}

func (s *serviceProvider) GetChatService(ctx context.Context) chatService.Service {
	if s.chatService != nil {
		return s.chatService
	}

	s.chatService = chatService.NewChatService(s.GetChatRepository(ctx))
	return s.chatService
}
