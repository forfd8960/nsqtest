package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	messageBody := []byte("hello")
	topic := "test-topic"
	if err := producer.Publish(topic, messageBody); err != nil {
		log.Fatal(err)
	}

	producer.Stop()
}
