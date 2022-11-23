package main

import (
	"github.com/MMMahmoodian/alarm/conf"
	"github.com/MMMahmoodian/alarm/conf/rabbit_mq"
	"github.com/MMMahmoodian/alarm/models"
	"github.com/MMMahmoodian/alarm/services/queue"
	"log"
)

func init() {
	conf.Initialize()
}

func main() {
	messageModel := models.TelegramMessage{}
	queueName := rabbit_mq.TelegramQueue
	consumer := queue.RMQConsumer{}
	err := consumer.Consume(messageModel, queueName)
	if err != nil {
		log.Fatal(err)
	}
}
