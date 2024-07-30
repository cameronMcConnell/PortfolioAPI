package main

import (
	"log"
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

	err = server.StartServer()
	if err != nil {
		log.Fatal(err)
	}
}