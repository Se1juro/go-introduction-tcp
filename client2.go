package main

import (
	"encoding/gob"
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
	err := gob.NewEncoder(connection).Encode(&message)
	if err != nil {
		fmt.Println(err)
	}
	connection.Close()
}

func main() {
	go client()

	var input string
	fmt.Scanln(&input)
}
