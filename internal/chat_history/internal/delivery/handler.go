package delivery

import (
	"chat/internal"
	"chat/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type ChatHandler struct {
	Usecase internal.ChatUseCase
}

func NewChatHandler(usecase internal.ChatUseCase) *ChatHandler {
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

//TODO распарсить куку авторизации и добавить параметр get chat_id
func (c *ChatHandler) GetChat(w http.ResponseWriter, r *http.Request) {

	chat := models.Chat{}
	chat_id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err == nil {
		http.NotFound(w, r)
		return
	}

	Conv := models.Conversation{
		ConversationId: chat_id,
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
	uuid = r.Header.Get("Unparsed") //TODO придумать как аккуратно передавать пользователя которому нужен запрос

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
