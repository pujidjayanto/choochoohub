package main

import (
	"log"

	"github.com/pujidjayanto/choochoohub/user-api/bootstrap"
)

func main() {
	app, err := bootstrap.InitApp()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}
}
