package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	"our-little-chatik/internal/chat"
	"our-little-chatik/internal/models"
)

type ChatHandler struct {
	Usecase chat.ChatUseCase
}

func NewChatHandler(usecase chat.ChatUseCase) *ChatHandler {
	return &ChatHandler{
		Usecase: usecase,
	}
}

func (c *ChatHandler) PostMessage(w http.ResponseWriter, r *http.Request) {
	mes := models.Message{}

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&mes)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := c.Usecase.SaveMessage(mes); err != nil { // добавить контекст
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	w.WriteHeader(http.StatusOK)
}

func (c *ChatHandler) GetChat(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	chat := models.Chat{}
	err := json.NewDecoder(r.Body).Decode(&chat)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	Conv := models.Conversation{
		ConversationId: chat.ConversationId,
		Owner:          chat.Owner,
	}

	if Conv.MessageList, err = c.Usecase.FetchChat(chat); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	resp, err := json.Marshal(&Conv)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(resp)

	w.WriteHeader(http.StatusOK)
}

func (c *ChatHandler) GetChatList(w http.ResponseWriter, r *http.Request) {

	var uuid string
	if err := json.NewDecoder(r.Body).Decode(&uuid); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	List, err := c.Usecase.ChatList(uuid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	result := models.ChatList{Owner: uuid, List: List}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}

	w.WriteHeader(http.StatusOK)
}
