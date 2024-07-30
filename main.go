package main

import (
	"log"
	"fmt"
	"github.com/cameronMcConnell/PortfolioAPI/lib"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	serverAddress := lib.ReadEnv("SERVER_ADDRESS")

	server := lib.NewServer(serverAddress)

	fmt.Printf("Starting server on: %s\n", serverAddress)

	err = server.StartServer()
	if err != nil {
		log.Fatal(err)
	}
}