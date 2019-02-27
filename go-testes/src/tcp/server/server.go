package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]

	fmt.Println("Launching server...")

	ln, err := net.Listen("tcp", PORT)

	if err != nil {
		fmt.Println(err)
		return
	}

	for {

		conn, err := ln.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}

		//handleConnection(conn)

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(message))
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))

		conn.Close()
	}
}

// func handleConnection(conn net.Conn) {

// 	conn.Close()
// }
