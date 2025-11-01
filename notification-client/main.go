package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "otp.created",
		GroupID: "otp-consumer-group",
	})

	defer reader.Close() //nolint:errcheck

	fmt.Println("Listening for messages on topic otp.created...")

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received message: %s\n", string(m.Value))
	}
}
