package main

import (
	"myapp/internal/config"
	"myapp/internal/server"
	"log"
)

func main() {
	log.Println("Starting myapp...")
	
	// Load application configuration
	cfg := config.LoadConfig()
	
	// Create and start the server
	s := server.NewServer(cfg)
	s.Start()
}