package main

import (
	"fmt"
	"net"
	"strings"
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
		}

		if value.typ != "array" {
			fmt.Println("Invalid request, expected array")
			continue
		}

		if len(value.array) == 0 {
			fmt.Println("Invalid request, expected array length > 0")
			continue
		}

		fmt.Printf("%+v\n", value)

		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		writer := NewWriter(connection)
		fmt.Println("commands are ", command)
		handler, ok := Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			writer.Write(Value{typ: "string", str: ""})
			continue
		}

		result := handler(args)
		writer.Write(result)
	}
}

func main() {
	setupRedisServer()
}
