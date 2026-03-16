package bybit_adapter

import (
	"context"
	"errors"
	"investment-dashboard/interfaces"
	"strconv"
)

func (adapter *BybitAdapter) getEarns(ctx context.Context, result interfaces.Result) error {
	params := map[string]interface{}{"accountType": "FUND", "category": "FlexibleSaving"}
	balance, err := adapter.Client.NewClassicalBybitServiceWithParams(params).GetEarnRedeemPosition(ctx)

	if err != nil {
		return err
	}

	return adapter.formatEarns(ctx, balance.Result.(map[string]interface{})["list"], result)
}

func (adapter *BybitAdapter) formatEarns(ctx context.Context, data interface{}, result interfaces.Result) error {
	rawEarnsData, ok := data.([]interface{})
	if !ok {
		// TODO: make norma handle of that error
		return errors.New("invalid data from bybit endpoint")
	}

	for _, value := range rawEarnsData {

		var coin interfaces.Position

		if coinData, ok := value.(map[string]interface{})["coin"]; ok {
			currentCoinInResult, ok := result[coinData.(string)]

			if ok {
				coin = currentCoinInResult
			} else {
				coin.Name = coinData.(string)
			}
		}

		if amount, ok := value.(map[string]interface{})["amount"]; ok {
			quantity, _ := strconv.ParseFloat(amount.(string), 64)
			coin.EarnQuantity = quantity
			coin.Quantity += quantity
		}
		if coin.Quantity > 0 {
			result[coin.Name] = coin
		}
	}

	return nil
}
