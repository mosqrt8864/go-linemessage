package domain

import "context"

type Message struct {
	ID     string `bson:"id"`      //Message ID
	UserID string `bson:"user_id"` //User ID
	Text   string `bson:"text"`    //Message text
}
type MessageRepository interface {
	Add(ctx context.Context, msg Message) error
	GetUserMessages(ctx context.Context, userId string) ([]Message, error)
}
