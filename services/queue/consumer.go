package queue

import "github.com/MMMahmoodian/alarm/models"

type Consumer interface {
	Consume(messageModel models.Message, queue string) error
}
