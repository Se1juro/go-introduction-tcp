package main

import (
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
	data := make([]byte, 100) // Get data from client
	bytesRead, err := client.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Message: ", string(data[0:bytesRead]))
		fmt.Println("Bytes: ", bytesRead)
	}
}

func main() {
	go server() // Listen connections

	var input string
	fmt.Scanln(&input)
}
