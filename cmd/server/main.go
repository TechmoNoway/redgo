package main

import (
	"log"

	"github.com/TechmoNovay/redgo/internal/server"
)

func main() {
	srv := server.New(":6379")

	log.Println("RedGo listening on :6379")

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
