package delivery

import (
	"linemessage/pkg/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QueryMessageResp struct {
	UserID string `json:"user_id"` //User ID
	ID     string `json:"id"`      //Message ID
	Text   string `json:"text"`    //Message text
}

type QueryMessageReq struct {
	UserID string `json:"user_id"` //User ID
}

type SendMessageReq struct {
	UserID string `json:"user_id"` // User ID
	Text   string `json:"text"`    //Message text
}

type MessageHandler struct {
	messageUsecase domain.MessageUsecase
}

func NewMessageHandler(e *gin.Engine, messageUserCase domain.MessageUsecase) {
	handler := &MessageHandler{
		messageUsecase: messageUserCase,
	}
	e.POST("/api/v1/webhooks", handler.Webhooks)
	e.POST("/api/v1/messages", handler.SendMessage)
	e.GET("/api/v1/messages", handler.QueryMessage)
}

func (h *MessageHandler) Webhooks(c *gin.Context) {
	if err := h.messageUsecase.Webhooks(c, c.Request); err != nil {
		log.Println("Webhooks UseCase Error")
		c.JSON(http.StatusInternalServerError, "Webhooks UseCase Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

func (h *MessageHandler) SendMessage(c *gin.Context) {
	var request SendMessageReq
	if err := c.BindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
	}
	if err := h.messageUsecase.SendMessage(c, request.UserID, request.Text); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "SendMessage UseCase Error")
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

func (m *MessageHandler) QueryMessage(c *gin.Context) {
	var request QueryMessageReq
	request.UserID = c.Query("user_id")
	if request.UserID == "" {
		c.JSON(http.StatusBadRequest, "parameter error")
		return
	}
	messages, err := m.messageUsecase.QueryMessage(c, request.UserID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	// po to dto
	var response []QueryMessageResp
	for _, msg := range messages {
		response = append(response, QueryMessageResp{
			ID:     msg.ID,
			Text:   msg.Text,
			UserID: msg.UserID})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": response,
	})
}
