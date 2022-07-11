package main

import (
	"fmt"
	"net"
)

func client() {
	connection, error := net.Dial("tcp", ":3035")
	if error != nil {
		fmt.Println(error)
		return
	}
	message := "Connected to server"
	fmt.Println(message)
	connection.Write([]byte(message)) // Convert our message string to bytes, TCP only accepts bytes
	connection.Close()
}

func main() {
	go client()

	var input string
	fmt.Scanln(&input)
}
