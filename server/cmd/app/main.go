package main

import (
	"context"
	"fmt"
	bybit_adapter "investment-dashboard/infrastructure/adapters/bybit"
	"investment-dashboard/infrastructure/config"
)

func main() {
	fmt.Println("hello investment dashboard")
	ctx := context.Background()

	err := config.Config.Bootstrap()

	if err != nil {
		panic(err)
	}

	// it is mock, later move from env config to db store
	var MOCK_BYBIT_API_KEY, _ = config.Config.BYBIT_API_KEY()
	var MOCK_BYBIT_SECRET_KEY, _ = config.Config.BYBIT_SECRET_KEY()

	options := bybit_adapter.ConstructorOptions{
		ApiKey:    MOCK_BYBIT_API_KEY,
		SecretKey: MOCK_BYBIT_SECRET_KEY,
	}

	adapter := bybit_adapter.Constructor(options)

	adapter.GetEarns(ctx)
}
