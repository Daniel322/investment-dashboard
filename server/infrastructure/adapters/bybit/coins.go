package bybit_adapter

import (
	"context"
	"errors"
	"investment-dashboard/interfaces"
	"strconv"
)

func (adapter *BybitAdapter) getCoinsBalance(ctx context.Context, result interfaces.Result) error {
	params := map[string]interface{}{"accountType": "FUND", "category": "linear"}
	balance, err := adapter.Client.NewClassicalBybitServiceWithParams(params).GetAllCoinsBalance(ctx)

	if err != nil {
		return err
	}

	return adapter.formatCoins(ctx, balance.Result.(map[string]interface{})["balance"], result)
}

func (adapter *BybitAdapter) formatCoins(ctx context.Context, data interface{}, result interfaces.Result) error {
	rawCoinsData, ok := data.([]interface{})
	if !ok {
		// TODO: make norma handle of that error
		return errors.New("invalid data from bybit endpoint")
	}

	for _, value := range rawCoinsData {

		coin := interfaces.Position{}

		if coinData, ok := value.(map[string]interface{})["coin"]; ok {
			coin.Name = coinData.(string)
		}
		if walletBalance, ok := value.(map[string]interface{})["walletBalance"]; ok {
			quantity, _ := strconv.ParseFloat(walletBalance.(string), 64)
			coin.WalletQuantity = quantity
			coin.Quantity += quantity
		}
		if coin.Quantity > 0 {
			result[coin.Name] = coin
		}
	}

	return nil
}
