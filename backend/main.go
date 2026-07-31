package main

import (
	"log"

	"battleship/backend/internal/api"
)

func main() {
	server := api.NewServer()
	log.Println("Battleship backend listening on :8080")
	if err := server.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
