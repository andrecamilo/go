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

	defer ln.Close()

	for {

		conn, err := ln.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {

	i := 0

	for {

		message, _ := bufio.NewReader(conn).ReadString('\n')

		if message == "" {
			i++
		} else {
			i = 0
		}

		fmt.Print("Message Received:", string(message))
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))

		if i > 3 {
			break
		}
	}

	conn.Close()
}
