package main

import (
	"log"

	"github.com/timetravel-1010/indexer-api/internal/server"
)

func main() {

	server := server.NewServer()

	log.Println("starting server...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("cannot start server: %v", err.Error())
	}
}
