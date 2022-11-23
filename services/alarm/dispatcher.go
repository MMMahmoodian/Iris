package alarm

import (
	"github.com/MMMahmoodian/alarm/models"
	"github.com/MMMahmoodian/alarm/services/queue"
)

type Dispatcher struct{}

func (s Dispatcher) Dispatch(publisher queue.Publisher, message models.Message, queue string) error {
	if err := publisher.Publish(message, queue); err != nil {
		return err
	}
	return nil
}
