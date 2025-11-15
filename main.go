package main

import (
	"fmt"
	"io"
	"net"
	"os"
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
	}
	defer connection.Close()

	for {
		buffer := make([]byte, 1024)

		_, err := connection.Read(buffer)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading from client: ", err.Error())
			os.Exit(1)
		}
		connection.Write([]byte("+Ok\r\n"))
	}
}

func main() {
	setupRedisServer()
}
