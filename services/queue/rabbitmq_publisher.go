package queue

import (
	"context"
	"github.com/MMMahmoodian/alarm/conf/rabbit_mq"
	"github.com/MMMahmoodian/alarm/models"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type RMQPublisher struct{}

func (RMQPublisher) Publish(message models.Message, queue string) error {
	conn, err := rabbit_mq.Connection()
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := rabbit_mq.Channel(conn)
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := rabbit_mq.Queue(ch, queue)
	// Failed to declare a queue
	if err != nil {
		return err
	}

	body, err := message.ToBytes()
	// Failed to serialize
	if err != nil {
		return err
	}
	err = ch.PublishWithContext(context.Background(),
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	// Failed to publish a message
	if err != nil {
		return err
	}
	log.Println("sent!")
	return nil
}
