package infra

import "github.com/line/line-bot-sdk-go/v7/linebot"

func NewBot(secret string, token string) (*linebot.Client, error) {
	bot, err := linebot.New(secret, token)
	if err != nil {
		return nil, err
	}
	return bot, err
}
