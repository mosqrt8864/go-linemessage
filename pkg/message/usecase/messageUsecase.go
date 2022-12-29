package usecase

import (
	"context"
	"linemessage/pkg/domain"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type messageUsecase struct {
	messageRepo domain.MessageRepository
}

func NewMessageUsecase(messageRepo domain.MessageRepository) domain.MessageUsecase {
	return &messageUsecase{
		messageRepo: messageRepo,
	}
}

func (m *messageUsecase) Webhooks(ctx context.Context, events []*linebot.Event) error {
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				err := m.messageRepo.Add(ctx, domain.Message{
					ID:     message.ID,
					Text:   message.Text,
					UserID: event.Source.UserID,
				})
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
