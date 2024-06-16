package main

import (
	"log"
	"log-client-go/router"
)

func main() {
	r := router.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
