package main

import (
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

type messageHandler struct{}

func (h *messageHandler) HandleMessage(message *nsq.Message) error {
	fmt.Printf("received msg: %v\n", string(message.Body))
	return nil
}

func main() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test-topic", "test-chan", config)
	if err != nil {
		fmt.Println("newConsumer error: ", err)
		return
	}

	consumer.AddHandler(&messageHandler{})
	err = consumer.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		fmt.Println("connect nsqlookoupd error: ", err)
		return
	}

	time.Sleep(1 * time.Second)

	consumer.Stop()
}
