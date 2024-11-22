package main

import (
	"log"
	"net/http"

	"image_resampler/internal/api"
	"image_resampler/internal/config"
)

func main() {
	cfg := config.ParseFlags()

	router := api.NewRouter(cfg)
	log.Printf("Starting server on port 8085")
	if err := http.ListenAndServe(":8085", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
