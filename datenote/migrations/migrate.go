package main

import (
	"datenote/datenote/storage/postgres"
	"log"
)

func main() {
	if err := postgres.Migrate(); err != nil {
		log.Fatal(err)
	}
}
