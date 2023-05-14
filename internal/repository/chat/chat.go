package chat

import (
	"github.com/Slintox/chat_server/pkg/database/postgres"
)

const tableName = `chat`

type Repository interface {
	// Methods
}

type repository struct {
	client postgres.Client
}

func NewRepository(client postgres.Client) Repository {
	return &repository{
		client: client,
	}
}
