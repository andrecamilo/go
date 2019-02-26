package main

import (
	"fmt"
	"time"
)

func falar(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main1() {
	go falar("ola")
	go falar("tudo bem")
}
