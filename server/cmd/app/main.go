package main

import (
	"investment-dashboard/infrastructure/config"
	"investment-dashboard/infrastructure/database"
	"log"
)

func main() {
	log.Println("init investment dashboard")

	err := config.Config.Bootstrap()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
