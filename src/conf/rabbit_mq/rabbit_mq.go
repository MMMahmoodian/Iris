package rabbit_mq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
)

const (
	TelegramQueue string = "RABBIT_MQ_TELEGRAM_QUEUE"
)

func Connection() (*amqp.Connection, error) {
	user := os.Getenv("RABBITMQ_USER")
	pw := os.Getenv("RABBITMQ_PASSWORD")
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")
	connection := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pw, host, port)
	return amqp.Dial(connection)
}

func Channel(conn *amqp.Connection) (*amqp.Channel, error) {
	return conn.Channel()
}

func Queue(ch *amqp.Channel, queue string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		os.Getenv(queue),
		false,
		false,
		false,
		false,
		nil,
	)
}
