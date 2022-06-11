package internal

import (
	"our-little-chatik/internal/peer/internal/models"
)

type WebSocketWorker interface {
	Read()
	Write()
	Close()
}

type PeerUsecase interface {
	SendMessage(msg *models.Message, chat *models.Chat) error
	FetchMessages(chat *models.Chat) ([]models.Message, error)
}

type PeerRepo interface {
	SendPayload(msg *models.Message, chat *models.Chat) error
	FetchUpdates(chat *models.Chat) ([]models.Message, error)
}

type MessageManager interface {
	Work()
	EnqueueChatIfNotExists(msg *models.Message) *models.Chat
}
