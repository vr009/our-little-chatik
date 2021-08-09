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

func SaveMessage(mes models.Message) error {
	return nil
}
func FetchChat(chat models.Chat) ([]models.Message, error) {
	return []models.Message{}, nil
}
func ChatList(userId string) ([]models.Chat, error) {
	return []models.Chat{}, nil
}
