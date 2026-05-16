package main

import (
	"guardian-grid-agent/internal/app"
	"os"
)

func main() {
	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://127.0.0.1:8080"
	}
	app.RunAgent(apiURL)
}
