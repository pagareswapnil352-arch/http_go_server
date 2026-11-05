package main

import (
	"errors"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {

		panic(errors.New("failed to listen on specified address"))
	}

	conn, err := listener.Accept()

	if err != nil {

		panic(errors.New("failed to accept connection"))

	}

	conn.Write([]byte("Hello, client!"))

	conn.Close()

}
