package domain

import (
	"context"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Message struct {
	ID     string `bson:"id"`      //Message ID
	UserID string `bson:"user_id"` //User ID
	Text   string `bson:"text"`    //Message text
}
type MessageRepository interface {
	Add(ctx context.Context, msg Message) error
	GetUserMessages(ctx context.Context, userId string) ([]Message, error)
}

type MessageUsecase interface {
	Webhooks(context.Context, []*linebot.Event) error
}
