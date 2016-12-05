package main

import "net"

func main() {
	listener, err := net.Listen("tcp", ":7654")
	if err != nil {
		panic(err)
	}

	for {
		_, err := listener.Accept()
		if err != nil {
			panic(err)
		}
	}
}
