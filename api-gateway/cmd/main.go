package main

import (
	api "api-gateway/internal/api/server"
	"log"
)

func main() {
	log.Println("Starting API Gateway... 🌐")
	api.RunServer()
}
