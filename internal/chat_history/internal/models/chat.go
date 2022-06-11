package models

import "time"

type Chat struct {
	ConversationId int       `json:"conversation_id" bson:"conversation_id"`
	Owner          string    `json:"owner" bson:"owner"`
	Opponent       string    `json:"opponent" bson:"opponent"`
	LastMessage    time.Time `json:"last_message" bson:"last_message"`
}

type ChatList struct {
	Owner string `json:"owner" bson:"owner"`
	List  []Chat `json:"list" bson:"list"`
}

type Conversation struct {
	ConversationId int    `json:"conversation_id" bson:"conversation_id"`
	Owner          string `json:"owner" bson:"owner"`
	MessageList    []Message
}
