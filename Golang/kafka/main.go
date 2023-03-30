package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	kafkago "github.com/segmentio/kafka-go"
)

func main() {
	writer := kafkago.NewWriter(kafkago.WriterConfig{
		Brokers: []string{"shared-bootstrap.infra.samarkand.io:9094"},
		Topic:   "test-topic",
	})
	defer writer.Close()
	val := map[string]string{"test": "val"}
	bval, _ := json.Marshal(val)
	message := kafkago.Message{
		Key:   []byte("key"),
		Value: bval,
	}

	// create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// write the message to the Kafka topic with the timeout context
	err := writer.WriteMessages(ctx, message)
	if err != nil {
		fmt.Println(err)
		// handle error
	}
}
