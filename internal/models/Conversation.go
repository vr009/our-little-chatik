package models

type Conversation struct {
	ConversationId string `json:"conversation_id"`
	Owner          string `json:"owner"`
	MessageList    []Message
}
