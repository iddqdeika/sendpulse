package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Logger interface {
	Log(message string)
	Alert(message string)
}

const logFileName = "log.log"

type DefaultLogger struct {
}

func (l *DefaultLogger) Log(message string) {
	fmt.Println(message)
}

func (l *DefaultLogger) Alert(message string) {
	fmt.Println(message)
	l.writeLog(message)
}

func (l *DefaultLogger) writeLog(message string) {
	date := time.Now()
	message = date.Format(time.StampMilli) + ": " + message

	err := ioutil.WriteFile(logFileName, []byte(message), os.ModeAppend)
	if err != nil {
		l.out(err.Error())
		l.out("Can't write log file, alert!")
	}
}

func (l *DefaultLogger) out(message string) {
	fmt.Println(message)
}
