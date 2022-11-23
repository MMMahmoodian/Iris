package alarm

import (
	"github.com/MMMahmoodian/alarm/models"
	"github.com/MMMahmoodian/alarm/services/queue"
)

type Handler struct{}

func (s Handler) Handle(consumer queue.Consumer, messageModel models.Message, queue string) error {
	if err := consumer.Consume(messageModel, queue); err != nil {
		return err
	}
	return nil
}
