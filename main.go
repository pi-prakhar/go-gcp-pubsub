package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	ctx := context.Background()

	// Create a client
	client, err := pubsub.NewClient(ctx, PROJECT_ID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Create a topic (or use an existing one)
	topic := client.Topic(TOPIC_ID)

	// Publish a message
	result := topic.Publish(ctx, &pubsub.Message{
		Data:       []byte("Hello, World!"),
		Attributes: map[string]string{"sub": "1"},
	})

	// Get the result of the publish call
	id, err := result.Get(ctx)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
	fmt.Printf("Published message with ID: %s\n", id)

}
