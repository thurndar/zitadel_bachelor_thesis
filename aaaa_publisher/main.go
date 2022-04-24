package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {

	nats_instance := "localhost:4222"
	subject := "foo"
	message := "Hello World!"

	// Connect to NATS
	nc, err := nats.Connect(nats_instance)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Simple Publisher
	nc.Publish(subject, []byte(message))

	nc.Flush()

}
