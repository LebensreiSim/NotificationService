package main

import (
	"log"

	"notifications/internal/app"
)

func main() {
	err := app.Start()
	if err != nil {
		log.Fatalf("failed to start the app: %v", err)
	}
}
