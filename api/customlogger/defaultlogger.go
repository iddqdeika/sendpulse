package customlogger

import (
	"log"
)

type DefaultLogger struct{

}

func (l *DefaultLogger) Log(message string){
	log.Println(message)
}

func (l *DefaultLogger) Alert(err error){
	log.Println(err)
}