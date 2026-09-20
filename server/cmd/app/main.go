package main

import (
	"context"
	"log"
	"time"

	"investment-dashboard/infrastructure/config"
	"investment-dashboard/infrastructure/database"
	"investment-dashboard/internal/asset"
	"investment-dashboard/internal/events"
	"investment-dashboard/pkg/event_bus"
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

	eventBus := event_bus.NewBus()

	events.Init(db, eventBus)
	assetModule := asset.Init(db, eventBus)
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

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
	}
}
