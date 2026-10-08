package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/TechmoNovay/redgo/internal/command"
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

		cmd, err := command.Parse(message)
		if err != nil {
			conn.Write([]byte("ERR invalid command\n"))
			continue
		}

		fmt.Printf("command=%s args=%v\n", cmd.Name, cmd.Args)

		switch cmd.Name {
		case "PING":
			conn.Write([]byte("PONG\n"))

		default:
			conn.Write([]byte("ERR unknown command\n"))

		}
	}
}
