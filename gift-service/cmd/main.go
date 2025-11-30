package main

import (
	api "gift-service/internal/api/server"
	"log"
)

func main() {
	log.Println("Starting Gift Service... ✅")
	api.RunServer()
}
