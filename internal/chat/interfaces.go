package chat

import "our-little-chatik/internal/models"

type ChatRepo interface {
	AddMessage(mes models.Message) error
	GetChat(chat models.Chat) ([]models.Message, error)
	GetChatList(userId string) ([]models.Chat, error)
}

type ChatUseCase interface {
	SaveMessage(mes models.Message) error
	FetchChat(chat models.Chat) ([]models.Message, error)
	ChatList(userId string) ([]models.Chat, error)
}
