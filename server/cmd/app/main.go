package main

import (
	"context"
	"log"

	"investment-dashboard/infrastructure/config"
	"investment-dashboard/infrastructure/database"
	"investment-dashboard/internal/asset"
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

	assetModule := asset.Init(db)

	price := 0
	rate := 0
	res, err := assetModule.CreateAssetCommandHandler.Handle(context.Background(), asset.CreateAssetCommand{
		Name:  "Bitcoin",
		Slug:  "BTC",
		Type:  "crypto",
		Price: &price,
		Rate:  &rate,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
