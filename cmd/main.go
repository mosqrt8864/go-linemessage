package main

import (
	"context"
	"fmt"
	"linemessage/pkg/infra"
	"linemessage/pkg/message/delivery"
	"linemessage/pkg/message/repository"
	"linemessage/pkg/message/usecase"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	ctx := context.Background()
	log.Println("LineMessage application start")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// get config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("config yaml Error")
		return
	}
	channelSecret := viper.GetString("line.channel-secret")
	channelAccessToken := viper.GetString("line.channel-access-token")
	appPort := viper.GetInt("application.port")
	dbUserName := viper.GetString("mongodb.username")
	dbPassword := viper.GetString("mongodb.password")
	dbAddress := viper.GetString("mongodb.address")
	dbPort := viper.GetInt("mongodb.port")

	conn, err := infra.ConnMongoDB(ctx, dbUserName, dbPassword, dbAddress, dbPort)
	if err != nil {
		log.Println("Please check mongodb is startup")
		return
	}
	lineBot, err := infra.NewBot(channelSecret, channelAccessToken)
	if err != nil {
		log.Println("LineBot Error")
		return
	}

	r := gin.Default()

	mongoDB := conn.Database("linemessage")
	messageRepo := repository.NewMongoMessageRepository(mongoDB)
	messageUseCase := usecase.NewMessageUsecase(messageRepo)
	delivery.NewMessageHandler(r, messageUseCase, lineBot)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", appPort),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server Error", err)
		}
	}()

	<-interrupt
	log.Println("LineMessage application end")
}
