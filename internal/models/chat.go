package models

import "time"

type Chat struct {
	ConversationId string    `json:"conversation_id"`
	Owner          string    `json:"owner"`
	Opponent       string    `json:"opponent"`
	LastMessage    time.Time `json:"last_message"`
}

type ChatList struct {
	Owner string `json:"owner"`
	List  []Chat `json:"list"`
}

type Conversation struct {
	ConversationId string `json:"conversation_id"`
	Owner          string `json:"owner"`
	MessageList    []Message
}
