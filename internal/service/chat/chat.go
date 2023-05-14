package chat

import (
	chatRepo "github.com/Slintox/chat_server/internal/repository/chat"
)

type Service interface {
	// Methods
}

type chatService struct {
	chatRepo chatRepo.Repository
}

func NewChatService(chatRepo chatRepo.Repository) Service {
	return &chatService{
		chatRepo: chatRepo,
	}
}
