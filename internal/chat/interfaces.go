package chat

import (
	models2 "chat/internal/models"
)

type ChatRepo interface {
	AddMessage(mes models2.Message) error
	GetChat(chat models2.Chat) ([]models2.Message, error)
	GetChatList(userId string) ([]models2.Chat, error)
}

type ChatUseCase interface {
	SaveMessage(mes models2.Message) error
	FetchChat(chat models2.Chat) ([]models2.Message, error)
	ChatList(userId string) ([]models2.Chat, error)
}
