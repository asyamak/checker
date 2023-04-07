package main

import (
	"log"

	"restapi/config"
	"restapi/internal/app"
)

const (
	filePath = "./config/config"
)

func main() {
	cnf, err := config.New(filePath)
	if err != nil {
		log.Fatalf("error loading config")
	}

	err = app.Start(cnf)
	if err != nil {
		log.Fatalf("app: error initiate application")
	}
}
