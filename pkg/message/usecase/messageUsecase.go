package usecase

import (
	"context"
	"linemessage/pkg/domain"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type messageUsecase struct {
	messageRepo domain.MessageRepository
	lineBot     *linebot.Client
}

func NewMessageUsecase(messageRepo domain.MessageRepository, lineBot *linebot.Client) domain.MessageUsecase {
	return &messageUsecase{
		messageRepo: messageRepo,
		lineBot:     lineBot,
	}
}

func (m *messageUsecase) Webhooks(ctx context.Context, req *http.Request) error {
	events, err := m.lineBot.ParseRequest(req)
	if err != nil {
		return err
	}
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

func (m *messageUsecase) SendMessage(ctx context.Context, userID string, text string) error {
	lineText := linebot.NewTextMessage(text)
	var messages []linebot.SendingMessage
	messages = append(messages, lineText)
	_, err := m.lineBot.PushMessage(userID, messages...).Do()
	if err != nil {
		return err
	}
	return nil
}
