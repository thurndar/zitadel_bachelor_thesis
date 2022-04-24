package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'", i, m.Subject, string(m.Data))
}

func main() {

	nats_instance := "localhost:4222"
	subscription := "foo"
	amountOfReceivedEvents := 0
	amountOfEventsToReceive := 1

	// Connect to NATS
	nc, err := nats.Connect(nats_instance)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Simple Subscriber
	log.Printf("Listening on [%s]", subscription)
	for amountOfReceivedEvents <= amountOfEventsToReceive {
		nc.Subscribe(subscription, func(msg *nats.Msg) {
			amountOfReceivedEvents += 1
			printMsg(msg, amountOfReceivedEvents)
		})
		nc.Flush()
	}
}
