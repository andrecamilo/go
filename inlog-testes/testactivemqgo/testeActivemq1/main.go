package main

import (
	"fmt"
	"testeActivemq1/activemq"
)

func main() {
	//Send
	//if err := activeMQ.NewActiveMQ("192.168.229.126:61613").Send("systrac.Bridge.TempoReal.dbDispBusNet", "test from 1"); err != nil {
	//fmt.Println("AMQ ERROR:", err)
	//}

	//this func will handle the messges get from activeMQ server.
	fmt.Println("inicio")

	if err := activemq.NewActiveMQ("127.0.0.1:61623").
		//if err := activemq.NewActiveMQ("failover:(tcp://127.0.0.1:61626)").
		Subscribe("/topic/systrac.Bridge.TempoReal.dbDispBusNet", handler); err != nil {
		//Subscribe("/queue/test-1", handler); err != nil {
		fmt.Println("AMQ ERROR:", err)
	}

	Send(destination string, msg string) error {

	fmt.Println("fim")
}

func handler(err error, msg string) {
	fmt.Println("chegou")
	fmt.Println("AMQ MSG:", err, msg)
}
