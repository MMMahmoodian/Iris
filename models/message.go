package models

type Message interface {
	Handle() error
	ToBytes() ([]byte, error)
	FromBytes(bytes []byte) (Message, error)
}
