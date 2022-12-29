package repository

import (
	"context"
	"linemessage/pkg/domain"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	messageCollection = "message"
)

type mongoMessageRepository struct {
	db *mongo.Database
}

func NewMongoMessageRepository(db *mongo.Database) domain.MessageRepository {
	return &mongoMessageRepository{
		db: db,
	}
}

func (m *mongoMessageRepository) Add(ctx context.Context, msg domain.Message) error {
	_, err := m.db.Collection(messageCollection).InsertOne(ctx, msg)
	if err != nil {
		return err
	}
	return nil
}

func (m *mongoMessageRepository) GetUserMessages(ctx context.Context, userId string) ([]domain.Message, error) {
	var msg domain.Message
	var messages []domain.Message
	cursor, err := m.db.Collection(messageCollection).Find(
		ctx,
		bson.D{{Key: "user_id", Value: userId}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		err := cursor.Decode(&msg)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
