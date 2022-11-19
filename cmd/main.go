package main

import (
	"log"

	"github.com/shortUrl/ICAssignment/config"
	"github.com/shortUrl/ICAssignment/internal/server"
)

func main() {
	log.Println("starting api server")
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	} else {
		log.Println("config loaded succesfully.")
	}

	log.Printf("AppVersion: %s \n", cfg.Server.AppVersion)

	srv := server.NewServer(cfg)

	srv.ListenAndServe()

}
