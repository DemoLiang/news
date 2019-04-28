package message

import (
	"errors"
	"news/db"
)

var messageChain chan db.Message

func Push(m db.Message) {
	messageChain <- m
}

func Pop() (db.Message, error) {
	select {
	case m := <-messageChain:
		return m, nil
	default:
		return db.Message{}, errors.New("nil")
	}
}

func init() {
	messageChain = make(chan db.Message, 10000)
}
