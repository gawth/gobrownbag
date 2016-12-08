package main

import (
	"io"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	io.Copy(os.Stdout, conn)
	conn.Close()
}
func main() {
	listener, err := net.Listen("tcp", ":7654")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}
