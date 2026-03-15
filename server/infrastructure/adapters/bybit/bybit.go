package bybit_adapter

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"

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

type CoinsData struct {
	bonus           string
	coin            string
	transferBalance string
	walletBalance   string
}

type Coin struct {
	Name     string
	Quantity float64
	Currency float64
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

func (adapter *BybitAdapter) GetBalance(ctx context.Context) {
	coinsBalance, err := adapter.getCoinsBalance(ctx)

	if err != nil {
		adapter.Logger.Println("error on get coins", err.Error())
	}

	// earnsBalance, err := adapter.getEarns(ctx)

	// if err != nil {
	// 	adapter.Logger.Println("error on get earns", err.Error())
	// }

	fmt.Println("coins \n", coinsBalance)
	// fmt.Println("earns \n", earnsBalance)
}

func (adapter *BybitAdapter) getCoinsBalance(ctx context.Context) (interface{}, error) {
	params := map[string]interface{}{"accountType": "FUND", "category": "linear"}
	balance, err := adapter.Client.NewClassicalBybitServiceWithParams(params).GetAllCoinsBalance(ctx)

	if err != nil {
		return nil, err
	}

	fmt.Println("RAW DATA \n", balance.Result.(map[string]interface{})["balance"], reflect.TypeOf(balance.Result.(map[string]interface{})["balance"]))

	coinsData := make([]CoinsData, 0)
	rawCoinsData, ok := balance.Result.(map[string]interface{})["balance"].([]interface{})
	if !ok {
		// TODO: make norma handle of that error
		return nil, errors.New("invalid data from bybit endpoint")
	}

	for _, value := range rawCoinsData {

		coinData := CoinsData{}

		if bonus, ok := value.(map[string]interface{})["bonus"]; ok {
			coinData.bonus = bonus.(string)
		}
		if coin, ok := value.(map[string]interface{})["coin"]; ok {
			coinData.coin = coin.(string)
		}
		if transferBalance, ok := value.(map[string]interface{})["transferBalance"]; ok {
			coinData.transferBalance = transferBalance.(string)
		}
		if walletBalance, ok := value.(map[string]interface{})["walletBalance"]; ok {
			coinData.walletBalance = walletBalance.(string)
		}

		coinsData = append(coinsData, coinData)
	}

	return adapter.formatCoins(ctx, coinsData), nil
}

func (adapter *BybitAdapter) formatCoins(ctx context.Context, data []CoinsData) []Coin {
	result := make([]Coin, 0)
	for _, coin := range data {
		quantity, _ := strconv.ParseFloat(coin.walletBalance, 64)
		result = append(result, Coin{Name: coin.coin, Quantity: quantity})
	}

	return result
}

func (adapter *BybitAdapter) getEarns(ctx context.Context) (interface{}, error) {
	params := map[string]interface{}{"accountType": "FUND", "category": "FlexibleSaving"}
	balance, err := adapter.Client.NewClassicalBybitServiceWithParams(params).GetEarnRedeemPosition(ctx)

	return balance.Result, err
}
