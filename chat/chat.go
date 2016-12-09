package main

import (
	"bufio"
	"fmt"
	"net"
)

type client struct {
	conn     net.Conn
	outgoing chan string
	incoming chan string
}
type chatRoom struct {
	clients []*client
	msgs    chan string
	joins   chan net.Conn
}

func (cl client) listen() {
	defer cl.conn.Close()
	reader := bufio.NewReader(cl.conn)
	for {
		line, _ := reader.ReadString('\n')
		fmt.Println("Got : ", line)
		cl.incoming <- line
	}
}
func (cl client) send() {
	writer := bufio.NewWriter(cl.conn)
	for data := range cl.outgoing {
		fmt.Println("Sending : ", data)
		writer.WriteString(data)
		writer.Flush()
	}
}
func (cr *chatRoom) handleConnection(conn net.Conn) {
	fmt.Println("Got a connection")
	cl := client{
		conn:     conn,
		outgoing: make(chan string),
		incoming: make(chan string),
	}
	cr.clients = append(cr.clients, &cl)
	go cl.listen()
	go cl.send()
	go func() {
		for {
			for data := range cl.incoming {
				fmt.Println("Piping data:", data)
				cr.msgs <- data
			}
		}
	}()
}

func (cr *chatRoom) broadCast(data string) {
	fmt.Println("Broadcasting:", data)
	for _, cl := range cr.clients {
		fmt.Println("    Broadcasting:", data)
		cl.outgoing <- data
	}
}
func (cr *chatRoom) coordinator() {
	go func() {
		for {
			select {
			case data := <-cr.msgs:
				cr.broadCast(data)
			case conn := <-cr.joins:
				cr.handleConnection(conn)
			}
		}
	}()
}
func main() {
	listener, err := net.Listen("tcp", ":7654")
	if err != nil {
		panic(err)
	}

	cr := chatRoom{
		clients: make([]*client, 0),
		msgs:    make(chan string),
		joins:   make(chan net.Conn, 0),
	}

	cr.coordinator()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		cr.joins <- conn
	}
}
