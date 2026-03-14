package bybit_adapter

import (
	"context"
	"fmt"
	"log"

	bybit_connector "github.com/bybit-exchange/bybit.go.api"
)

type BybitAdapter struct {
	Logger *log.Logger
	Client *bybit_connector.Client
}

type ConstructorOptions struct {
	ApiKey    string
	SecretKey string
}

// TODO: add formatter for requests data

func Constructor(options ConstructorOptions) *BybitAdapter {
	// init client
	fmt.Println(options)
	client := bybit_connector.NewBybitHttpClient(
		options.ApiKey,
		options.SecretKey,
		bybit_connector.WithBaseURL(bybit_connector.MAINNET),
		bybit_connector.WithDebug(true),
	)
	// return adapter pointer
	return &BybitAdapter{
		Client: client,
		Logger: client.Logger,
	}
}

func (adapter *BybitAdapter) GetCoinsBalance(ctx context.Context) (interface{}, error) {
	params := map[string]interface{}{"accountType": "FUND", "category": "linear"}
	balance, err := adapter.Client.NewClassicalBybitServiceWithParams(params).GetAllCoinsBalance(ctx)

	return balance.Result, err
}

func (adapter *BybitAdapter) GetEarns(ctx context.Context) (interface{}, error) {
	params := map[string]interface{}{"accountType": "FUND", "category": "FlexibleSaving"}
	balance, err := adapter.Client.NewClassicalBybitServiceWithParams(params).GetEarnRedeemPosition(ctx)

	return balance.Result, err
}
