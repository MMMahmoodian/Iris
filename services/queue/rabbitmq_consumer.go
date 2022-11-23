package queue

import (
	"github.com/MMMahmoodian/alarm/conf/rabbit_mq"
	"github.com/MMMahmoodian/alarm/models"
	"log"
)

type RMQConsumer struct{}

func (RMQConsumer) Consume(messageModel models.Message, queue string) error {
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	// Failed to register a consumer
	if err != nil {
		return err
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			message, err := messageModel.FromBytes(d.Body)
			if err != nil {
				log.Println("Failed to parse the message")
				d.Reject(false)
			}
			log.Printf("sending message to telegram")
			err = message.Handle()
			if err != nil {
				d.Reject(!d.Redelivered)
			} else {
				d.Ack(false)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
	return nil
}
