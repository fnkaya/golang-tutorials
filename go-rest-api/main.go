package main

import (
	"log"
	"wallet-service/server"
)

func main() {
	httpServer := server.NewServer(8080)
	err := httpServer.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
