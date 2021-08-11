package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"our-little-chatik/internal/models"
	"time"
)

type DataBase struct {
	clt mongo.Collection
	db  mongo.Database
}

func NewDataBase(ct mongo.Collection, db mongo.Database) *DataBase {
	return &DataBase{
		clt: ct,
		db:  db,
	}
}

func (db *DataBase) AddMessage(mes models.Message) error {
	chatId := mes.Sender + mes.Direction
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.clt.InsertOne(ctx, bson.D{{"_id", chatId}, {"name", "pi"}, {"value", 3.14159}})

	return err
}

func (db *DataBase) GetChat(chat models.Chat) ([]models.Message, error) {

	//out := []models.Message{}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := db.clt.Find(ctx, bson.D{{"_id", chat.ConversationId}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var episode bson.D
		if err = cur.Decode(&episode); err != nil {
			log.Fatal(err)
		}
		log.Println(episode)
		//var mes models.Message
		//err := bson.Unmarshal(episode,mes)

	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return nil, nil
}

func (db *DataBase) GetChatList(userId string) ([]models.Chat, error) {
	return nil, nil
}
