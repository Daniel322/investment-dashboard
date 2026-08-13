package main

import (
	"investment-dashboard/infrastructure/config"
	"log"
)

func main() {
	log.Println("init investment dashboard")

	err := config.Config.Bootstrap()

	if err != nil {
		log.Fatal(err)
	}
}
