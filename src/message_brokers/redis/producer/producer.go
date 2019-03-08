package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"../util"
)

func main() {

	svc := util.New(&util.NewInput{
		RedisURL: string("localhost:6379"), //os.Getenv("REDIS_URL"),
	})

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')

		err := svc.Publish("test/foo", text)
		if err != nil {
			log.Fatal(err)
		}
	}
}
