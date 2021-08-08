package chat

import "our-little-chatik/internal/models"

type ChatRepo interface {
	AddMessage(mes models.Message) error
	GetChat(chat models.Chat) error
	CreateChat(mes models.Message) error
	GetChatList(userId string) error
}

type ChatUseCase interface {
	SaveMessage(mes models.Message) error
	FetchChat(chat models.Chat) error
	ChatList(userId string) error
}
