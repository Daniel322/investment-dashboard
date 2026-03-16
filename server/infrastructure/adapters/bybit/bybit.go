package bybit_adapter

import (
	"context"
	"investment-dashboard/interfaces"

	bybit_connector "github.com/bybit-exchange/bybit.go.api"
)

func Constructor(options ConstructorOptions) *BybitAdapter {
	// init client
	client := bybit_connector.NewBybitHttpClient(
		options.ApiKey,
		options.SecretKey,
		bybit_connector.WithBaseURL(bybit_connector.MAINNET),
		bybit_connector.WithDebug(false),
	)
	// return adapter pointer
	return &BybitAdapter{
		Client: client,
		Logger: client.Logger,
	}
}

func (adapter *BybitAdapter) Balance(ctx context.Context) (interfaces.Result, error) {
	result := make(map[string]interfaces.Position)
	err := adapter.getCoinsBalance(ctx, result)

	if err != nil {
		adapter.Logger.Println("error on get coins", err.Error())
		return nil, err
	}

	err = adapter.getEarns(ctx, result)

	if err != nil {
		adapter.Logger.Println("error on get earns", err.Error())
		return nil, err
	}

	return result, nil
}
