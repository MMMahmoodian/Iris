package rabbit_mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
)

const (
	TelegramQueue string = "RABBIT_MQ_TELEGRAM_QUEUE"
)

func Connection() (*amqp.Connection, error) {
	return amqp.Dial(os.Getenv("RABBIT_MQ_CONNECTION"))
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
