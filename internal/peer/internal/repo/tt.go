package repo

import (
	"errors"
	"github.com/tarantool/go-tarantool"
	"log"
	"our-little-chatik/internal/peer/internal/models"
)

// TODO rewrite it to websockets
// See https://github.com/tarantool/websocket

type TarantoolRepo struct {
	conn *tarantool.Connection
}

func NewTarantoolRepo(conn *tarantool.Connection) *TarantoolRepo {
	return &TarantoolRepo{conn: conn}
}

func (tt *TarantoolRepo) SendPayload(msg *models.Message, chat *models.Chat) error {
	conn := tt.conn
	resp, err := conn.Call("put", []interface{}{
		chat.ChatID.String(),
		msg.SenderID.String(),
		msg.ReceiverID.String(),
		msg.Payload})
	if err != nil {
		return err
	}
	if resp == nil {
		return errors.New("Response is nil after Call")
	}
	if len(resp.Data) < 1 {
		return errors.New("Response.Data is empty after Eval")
	}
	return nil
}
func (tt *TarantoolRepo) FetchUpdates(chat *models.Chat) ([]models.Message, error) {
	var msgs []models.Message
	conn := tt.conn
	err := conn.CallTyped("take_msgs",
		[]interface{}{chat.ChatID.String(),
			chat.ReceiverID.String(),
			chat.SenderID.String(),
			chat.ReceiverID.String()}, &msgs)
	if err != nil && len(msgs) < 1 {
		log.Println("error from tt: ", err)
		return nil, err
	}

	if len(msgs) > 0 && msgs[0].Payload == "" {
		return nil, nil
	}

	return msgs, nil
}

func (tt *TarantoolRepo) Close() {
	tt.conn.Close()
}
