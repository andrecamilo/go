package main

import (
	"log"

	"../util"
)

func main() {

	svc := util.New(&util.NewInput{
		RedisURL: string("localhost:6379"), //os.Getenv("REDIS_URL"),
	})

	for {
		reply := make(chan []byte)
		err := svc.Subscribe("test/foo", reply)
		if err != nil {
			log.Fatal(err)
		}
		msg := <-reply
		print("Recieved: ", string(msg))
	}
}
