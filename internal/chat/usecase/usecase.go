package usecase

import (
	"our-little-chatik/internal/chat"
	"our-little-chatik/internal/models"
)

type ChatUseCase struct {
	repo chat.ChatRepo
}

func NewChatUseCase(rep chat.ChatRepo) *ChatUseCase {
	return &ChatUseCase{repo: rep}
}

func (ch *ChatUseCase) SaveMessage(mes models.Message) error {
	if ch.repo.ChatExist(mes.Sender + mes.Direction) {
		return ch.repo.AddMessage(mes)
	}
	err := ch.repo.CreateChat(mes)
	if err != nil {
		return err
	}

	return ch.repo.AddMessage(mes)
}
func (ch *ChatUseCase) FetchChat(chat models.Chat) ([]models.Message, error) {
	return ch.repo.GetChat(chat)
}
func (ch *ChatUseCase) ChatList(userId string) ([]models.Chat, error) {
	return ch.repo.GetChatList(userId)
}
