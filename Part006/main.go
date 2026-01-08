package main

import (
	"fmt"
	"os"
	"test/logger"
	"time"
)

type IReader interface {
	Read() string
}

type IWriter interface {
	Write(s string)
}

type LogMessage struct {
	PaymentId string
	UserId    string
	msg       string
	formated  string
}

func (l *LogMessage) String() string {
	l.formated = fmt.Sprintf("%s :  PaymentId: %s, UserId: %s : msg :%s", time.Now().String(), l.PaymentId, l.UserId, l.msg)
	return l.formated
}

func paymentBusinessLogic(logger logger.ILogger) {

	userId := "123"
	paymentId := "222"

	logger.Info(&LogMessage{paymentId, userId, "payment Created", ""})
	logger.Info(&LogMessage{paymentId, userId, "payment Updated", ""})
	logger.Fatal(&LogMessage{paymentId, userId, "payment Failed", ""})

	for _, s := range logger.ReadLogs() {
		fmt.Println(s)
	}
}

type Config struct {
	LogDriver string
}

func GetConfig() Config {
	return Config{
		LogDriver: os.Getenv("LOG_DRIVER"), // .env
	}
}

func main() {
	log := logger.NewLogger(GetConfig().LogDriver)
	paymentBusinessLogic(log)
}
