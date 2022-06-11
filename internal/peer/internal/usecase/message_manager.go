package usecase

import (
	"container/list"
	debug "log"
	"our-little-chatik/internal/peer/internal"
	"our-little-chatik/internal/peer/internal/models"
)

type MessageManagerImpl struct {
	chatList *list.List
	repo     internal.PeerRepo
}

func NewMessageManager(repo internal.PeerRepo) *MessageManagerImpl {
	chatList := list.New()
	return &MessageManagerImpl{repo: repo, chatList: chatList}
}

func (m *MessageManagerImpl) EnqueueChat(chat *models.Chat) {
	ptr := chat
	m.chatList.PushFront(ptr)
}

func (m *MessageManagerImpl) DequeueChat(chat *models.Chat) {
	for e := m.chatList.Front(); e != nil; e = e.Next() {
		if e.Value.(models.Chat).ChatID == chat.ChatID {
			m.chatList.Remove(e)
		}
	}
}

func (m *MessageManagerImpl) Work() {
	for {
		for e := m.chatList.Front(); e != nil; e = e.Next() {
			chat := e.Value.(*models.Chat)
			msgs, _ := m.repo.FetchUpdates(chat)
			select {
			case msg := <-chat.ReadyForSend:
				m.repo.SendPayload(msg, chat)
				debug.Println("sending: ", msg)
			default:
				if msgs != nil {
					chat.PutMsgsToRecv(msgs)
					debug.Println("some new messages here: ", msgs)
				}
			}
		}
	}
}

func (m *MessageManagerImpl) EnqueueChatIfNotExists(msg *models.Message) (chat *models.Chat) {
	for e := m.chatList.Front(); e != nil; e = e.Next() {
		*chat = e.Value.(models.Chat)
		if chat.ChatID == msg.ChatID {
			return chat
		}
	}
	chat = models.NewChatFromMsg(msg)
	m.EnqueueChat(chat)
	return chat
}
