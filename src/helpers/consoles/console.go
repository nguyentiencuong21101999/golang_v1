package c

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var defaultLogger = Logger{log.New(os.Stdout, "\r\n", 0)}

func (logger Logger) Print(values ...interface{}) {
	logger.Println(defaultLogger)
}

// Logger default logger
type Logger struct {
	LogWriter
}
type LogWriter interface {
	Println(v ...interface{})
}

func Log(m interface{}) {
	logs, err := json.MarshalIndent(m, "", "")
	if err != nil {
		fmt.Println("err : ", err)
	}
	fmt.Sprint(" %s\n", logs)
	log
}
