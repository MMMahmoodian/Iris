package queue

import "github.com/MMMahmoodian/alarm/models"

type Publisher interface {
	Publish(message models.Message, queue string) error
}
