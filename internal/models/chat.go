package models

import "time"

type Chat struct {
	ConversationId string    `json:"conversation_id"`
	Opponent       string    `json:"opponent"`
	LastMessage    time.Time `json:"last_message"`
}
