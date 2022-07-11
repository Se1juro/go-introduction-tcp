package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	listener, err := net.Listen("tcp", ":3035") // Create tcp server on 3035 port
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		newClient, err := listener.Accept() // Accept and return connection from new client
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(newClient)
	}
}

// What do we do with the client
func handleClient(client net.Conn) {
	var message string
	err := gob.NewDecoder(client).Decode(&message) // Convert bytes from client to string
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Message: ", message)
	}
}

func main() {
	go server() // Listen connections

	fmt.Println("Server run on port 3035")
	var input string
	fmt.Scanln(&input)
}
