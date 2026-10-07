package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Server struct {
	addr string
}

func New(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		fmt.Printf("client connected: %s\n", conn.RemoteAddr())

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("client disconnected: %s\n", conn.RemoteAddr())
			return
		}

		message = strings.TrimSpace(message)

		fmt.Printf("received: %s\n", message)

		if strings.EqualFold(message, "PING") {
			_, err := conn.Write([]byte("PONG\n"))
			if err != nil {
				return
			}
		} else {
			_, err := conn.Write([]byte("ERR unknown command\n"))
			if err != nil {
				return
			}
		}

	}

}
