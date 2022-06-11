package models

import (
	"github.com/google/uuid"
)

type Chat struct {
	ChatID       uuid.UUID `json:"chatID"`
	ReceiverID   uuid.UUID `json:"receiverID"`
	SenderID     uuid.UUID `json:"senderID"`
	ChatSibling  *Chat
	ReadyForRecv chan []Message
	ReadyForSend chan *Message
	msgForSend   *Message  // message to send to db
	msgsForRecv  []Message // messages to send to client
}

func (c *Chat) PutMsgToSend(m *Message) {
	c.ReadyForSend <- m
}

func (c *Chat) PutMsgsToRecv(m []Message) {
	c.ReadyForRecv <- m
}

func (c *Chat) FetchMsgsToRecv() []Message {
	return <-c.ReadyForRecv
}

func (c *Chat) waitForRecv() {
	<-c.ReadyForRecv
}

func (c *Chat) waitForSend() {
	<-c.ReadyForSend
}

func NewChatFromMsg(msg *Message) *Chat {
	chat := &Chat{
		ChatID:     msg.ChatID,
		ReceiverID: msg.ReceiverID,
		SenderID:   msg.SenderID,
	}

	chat.ReadyForRecv = make(chan []Message)
	chat.ReadyForSend = make(chan *Message)
	return chat
}
