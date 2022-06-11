package repo

import (
	models2 "chat/internal/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

type PGRepo struct {
	service             *sql.DB
	Db_name             string
	Chats_table_name    string
	Messages_table_name string
}

func NewPGRepo() *PGRepo {
	return &PGRepo{Db_name: "chats",
		Chats_table_name:    "Chats",
		Messages_table_name: "Messages",
	}
}

func (repo *PGRepo) InitDB() error {
	connStr := "host=db-chats port=5432 user=postgres password=admin dbname=chats sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	repo.service = db

	return nil
}

func (db *PGRepo) AddMessage(mes models2.Message) error {
	if db.service != nil {
		sender := mes.Sender

		//TODO оптимизировать либо проверку, либо перепроектировать
		var ok string
		checkstr := fmt.Sprintf("select from_id from chats where from_id='%s';", mes.Direction)
		result := db.service.QueryRow(checkstr)
		errq := result.Scan(&ok)
		if errq == nil {
			mes.Sender, mes.Direction = mes.Direction, mes.Sender
		}

		AddChatString := fmt.Sprintf("insert into %s values(default,'%s','%s') ON CONFLICT DO NOTHING;",
			db.Chats_table_name,
			mes.Sender,
			mes.Direction,
		)
		GetIdStr := fmt.Sprintf("select chat_id from %s where from_id='%s' and to_id='%s';", db.Chats_table_name, mes.Sender, mes.Direction)

		if _, err := db.service.Exec(AddChatString); err != nil {
			log.Print(err)
			return err
		}

		var chat_id int

		res := db.service.QueryRow(GetIdStr)
		err := res.Scan(&chat_id)
		if err != nil {
			log.Print(err)
			return err
		}

		AddMesString := fmt.Sprintf("insert into %s values(default,'%s',to_timestamp(%s),'%s','%s');",
			db.Messages_table_name,
			strconv.Itoa(chat_id),
			mes.Sent.Format("2014-04-04"),
			mes.Body,
			sender,
		)

		if _, err := db.service.Exec(AddMesString); err != nil {
			return err
		}

	}

	return nil
}

func (db *PGRepo) GetChat(chat models2.Chat) ([]models2.Message, error) {
	// TODO улучшить запрос на основании того что в структуре чат хранить id переписки для сокращения работы базы
	FetchString := fmt.Sprintf("select body, time from %s where chat_id=%s order by time desc;",
		db.Messages_table_name,
		strconv.Itoa(chat.ConversationId),
	)

	MessageList := make([]models2.Message, 0, 20)

	res, err := db.service.Query(FetchString)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer res.Close()
	for res.Next() {
		mesItem := models2.Message{}
		err := res.Scan(&mesItem.Body, &mesItem.Sent)
		if err != nil {
			log.Print(err)
		}
		MessageList = append(MessageList, mesItem)
	}

	return MessageList, nil
}

func (db *PGRepo) GetChatList(userId string) ([]models2.Chat, error) {

	//TODO добавить join для нормальной сортировки чатов
	ListString := fmt.Sprintf(
		"select chat_id, to_id from %s where from_id='%s' "+
			"union select chat_id, from_id from %s where to_id='%s';",
		db.Chats_table_name,
		userId,
		db.Chats_table_name,
		userId,
	)

	res, err := db.service.Query(ListString)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	chatList := make([]models2.Chat, 0, 5)
	for res.Next() {
		chatItem := models2.Chat{}
		err := res.Scan(&chatItem.ConversationId, &chatItem.Opponent)
		if err != nil {
			log.Print(err)
		}
		chatList = append(chatList, chatItem)
	}

	return chatList, nil
}

func (db *PGRepo) Close() {
	db.service.Close()
}
