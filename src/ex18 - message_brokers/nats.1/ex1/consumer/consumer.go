package main

import (
	"log"
	"time"

	"github.com/nats-io/go-nats"
)

func main() {

	//nc, err := nats.Connect("localhost:8222") //nats.DefaultURL)
	nc, err := nats.Connect("demo.nats.io") //nats.DefaultURL)

	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subscribe
	sub, err := nc.SubscribeSync("foo")
	if err != nil {
		log.Fatal(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Second)
		if err != nil {
			log.Fatal(err)
		}

		// Use the response
		log.Printf("Reply: %s", msg.Data)
	}

	// Close the connection
	nc.Close()
}
