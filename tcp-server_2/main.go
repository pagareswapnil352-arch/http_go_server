package main

import (
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {

		panic(err)
	}

	conn, err := listener.Accept()

	if err != nil {
		panic(err)
	}

	conn.Write([]byte("hii client"))

	conn.Close()

}
