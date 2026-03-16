package bybit_adapter

import (
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
