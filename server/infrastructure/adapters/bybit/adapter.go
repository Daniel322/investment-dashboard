package bybit_adapter

import (
	"encoding/json"
	"fmt"
	"investment-dashboard/interfaces"
	"io"
	"net/http"
	"strconv"
)

type BybitAdapter struct {
}

func New() *BybitAdapter {
	return &BybitAdapter{}
}

func (adapter *BybitAdapter) GetPrice(slug string) (*interfaces.GetPriceResponse, error) {
	response, err := http.Get(fmt.Sprintf("https://api.bybit.com/v5/market/tickers?category=spot&symbol=%s", adapter.slugToSymbol(slug)))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var getPriceResponse TickersResponse
	err = json.Unmarshal(body, &getPriceResponse)

	if err != nil {
		return nil, err
	}

	price, err := strconv.ParseFloat(getPriceResponse.Result.List[0].LastPrice, 64)

	if err != nil {
		return nil, err
	}

	return &interfaces.GetPriceResponse{
		Slug:  slug,
		Price: price,
	}, nil
}

func (adapter *BybitAdapter) slugToSymbol(slug string) string {
	return slug + "USDT"
}
