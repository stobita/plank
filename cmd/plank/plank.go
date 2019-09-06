package main

import (
	"log"

	"github.com/stobita/plank/internal/server"
)

func main() {
	log.Fatal(server.Run())
}
