package main

import (
	"log"
	api "wallet-service/internal/api/server"
)

func main() {
	log.Println("Starting Wallet Service...✅")
	api.RunServer()
}
