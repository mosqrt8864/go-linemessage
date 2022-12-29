package delivery

import (
	"linemessage/pkg/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type MessageHandler struct {
	messageUsecase domain.MessageUsecase
	lineBot        *linebot.Client
}

func NewMessageHandler(e *gin.Engine, messageUserCase domain.MessageUsecase, lineBot *linebot.Client) {
	handler := &MessageHandler{
		messageUsecase: messageUserCase,
		lineBot:        lineBot,
	}
	e.POST("/api/v1/webhooks", handler.Webhooks)
}

func (h *MessageHandler) Webhooks(c *gin.Context) {
	events, err := h.lineBot.ParseRequest(c.Request)
	if err != nil {
		log.Println("ParseRequest Error")
		c.JSON(http.StatusBadRequest, "ParseRequest Error")
		return
	}
	if err := h.messageUsecase.Webhooks(c, events); err != nil {
		log.Println("Webhooks UseCase Error")
		c.JSON(http.StatusInternalServerError, "Webhooks UseCase Error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}
