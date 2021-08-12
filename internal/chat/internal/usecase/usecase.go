package usecase

import (
	"chat"
	models2 "chat/internal/models"
)

type ChatUseCase struct {
	repo chat.ChatRepo
}

func NewChatUseCase(rep chat.ChatRepo) *ChatUseCase {
	return &ChatUseCase{repo: rep}
}

func (ch *ChatUseCase) SaveMessage(mes models2.Message) error {
	return ch.repo.AddMessage(mes)
}
func (ch *ChatUseCase) FetchChat(chat models2.Chat) ([]models2.Message, error) {
	return ch.repo.GetChat(chat)
}
func (ch *ChatUseCase) ChatList(userId string) ([]models2.Chat, error) {
	return ch.repo.GetChatList(userId)
}
