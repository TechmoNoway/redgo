package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Println("Connected to RedGo")

	_, err = conn.Write([]byte("SET name Kenny\n"))
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(conn)

	response, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server response: % s", response)

}
