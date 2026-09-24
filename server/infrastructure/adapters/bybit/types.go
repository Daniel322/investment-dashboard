package bybit_adapter

// TickersResponse — ответ от GET /v5/market/tickers
type TickersResponse struct {
	RetCode    int           `json:"retCode"`
	RetMsg     string        `json:"retMsg"`
	Result     TickersResult `json:"result"`
	RetExtInfo RetExtInfo    `json:"retExtInfo"`
	Time       int64         `json:"time"`
}

// RetExtInfo — расширенная информация об ошибке (обычно пустая при успехе)
type RetExtInfo struct {
	// Может содержать произвольные поля при ошибках, оставляем пустым
}

// TickersResult — результат запроса тикеров
type TickersResult struct {
	Category string   `json:"category"`
	List     []Ticker `json:"list"`
}

// Ticker — данные по одному торговому инструменту.
// В зависимости от category (spot/linear/inverse/option)
// некоторые поля могут отсутствовать или быть null.
type Ticker struct {
	Symbol                 string `json:"symbol"`
	LastPrice              string `json:"lastPrice"`
	IndexPrice             string `json:"indexPrice,omitempty"`
	MarkPrice              string `json:"markPrice,omitempty"`
	PrevPrice24h           string `json:"prevPrice24h"`
	Price24hPcnt           string `json:"price24hPcnt"`
	HighPrice24h           string `json:"highPrice24h"`
	LowPrice24h            string `json:"lowPrice24h"`
	PrevPrice1h            string `json:"prevPrice1h,omitempty"`
	OpenInterest           string `json:"openInterest,omitempty"`
	OpenInterestValue      string `json:"openInterestValue,omitempty"`
	Turnover24h            string `json:"turnover24h"`
	Volume24h              string `json:"volume24h"`
	FundingRate            string `json:"fundingRate,omitempty"`
	NextFundingTime        string `json:"nextFundingTime,omitempty"`
	PredictedDeliveryPrice string `json:"predictedDeliveryPrice,omitempty"`
	BasisRate              string `json:"basisRate,omitempty"`
	DeliveryFeeRate        string `json:"deliveryFeeRate,omitempty"`
	DeliveryTime           string `json:"deliveryTime,omitempty"`

	// Цены/объёмы лучших заявок
	Bid1Price string `json:"bid1Price"`
	Bid1Size  string `json:"bid1Size"`
	Ask1Price string `json:"ask1Price"`
	Ask1Size  string `json:"ask1Size"`

	// Только для опционов
	Bid1Iv          string `json:"bid1Iv,omitempty"`
	Ask1Iv          string `json:"ask1Iv,omitempty"`
	MarkIv          string `json:"markIv,omitempty"`
	UnderlyingPrice string `json:"underlyingPrice,omitempty"`

	// usdIndexPrice присутствует только для spot-пар с USD-индексом
	UsdIndexPrice string `json:"usdIndexPrice,omitempty"`
}
