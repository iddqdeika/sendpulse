package client

import (
	"log"
)

type DefaultLogger struct {
}

func (l *DefaultLogger) Log(message string) {
	log.Println(message)
}

func (l *DefaultLogger) Alert(message string) {
	log.Println(message)
}
