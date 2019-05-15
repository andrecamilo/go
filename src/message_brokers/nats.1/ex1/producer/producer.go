package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	nats "github.com/nats-io/go-nats"
)

func main() {

	//nc, err := nats.Connect("localhost:8222") //nats.DefaultURL)
	nc, err := nats.Connect("demo.nats.io") //nats.DefaultURL)

	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')

		if err := nc.Publish("foo", []byte(text)); err != nil {
			log.Fatal(err)
		}
	}

	// Make sure the message goes through before we close
	nc.Flush()
}
