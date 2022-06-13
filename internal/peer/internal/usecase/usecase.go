package usecase

import (
	"our-little-chatik/internal/peer/internal/models"
)

type PeerUsecaseImpl struct {
}

func NewPeerUsecaseImpl() *PeerUsecaseImpl {
	return &PeerUsecaseImpl{}
}

func (pu *PeerUsecaseImpl) SendMessage(msg *models.Message, chat *models.Chat) error {
	chat.PutMsgToSend(msg)
	return nil
}

func (pu *PeerUsecaseImpl) FetchMessages(chat *models.Chat) ([]models.Message, error) {
	msgs := chat.FetchMsgsToRecv()
	return msgs, nil
}
