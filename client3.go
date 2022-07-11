package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Person struct {
	Name  string
	Email []string
}

func client(person Person) {
	connection, error := net.Dial("tcp", ":3035")
	if error != nil {
		fmt.Println(error)
		return
	}
	message := "Connected to server"
	fmt.Println(message)
	err := gob.NewEncoder(connection).Encode(&person)
	if err != nil {
		fmt.Println(err)
	}
	connection.Close()
}

func main() {
	person := Person{
		Name: "Daniel",
		Email: []string{
			"daniel@gmail.com", "daniel2@gmail.com",
		},
	}
	go client(person)

	var input string
	fmt.Scanln(&input)
}
