package main

import (
	"github.com/harsh098/RPServer/internal/server"
	"log"
)

func main() {
	if err := server.RunServer(); err != nil {
		log.Fatalf("[RPServer] server run error: %v", err)
	}
}
