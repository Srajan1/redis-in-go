package main

import (
	"fmt"
	"net"
)

const (
	TYPE = "tcp"
	PORT = "6379"
	HOST = "localhost"
)

func setupRedisServer() {
	listener, error := net.Listen(TYPE, ":"+PORT)
	if error != nil {
		fmt.Println("Listening on TCP failed ", error)
	}
	fmt.Println("Redis Server running.")

	connection, error := listener.Accept()
	if error != nil {
		fmt.Println("Accepting connection failed ", error)
	} else {
		fmt.Println("Connection Accepted")
	}
	defer connection.Close()

	for {
		resp := NewResp(connection)
		// fmt.Println("New Resp object created")
		value, err := resp.Read()

		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("No errors observed during parsing input")
		}
		fmt.Println(value)
		connection.Write([]byte("+OK\r\n"))
	}
}

func main() {
	setupRedisServer()
}
