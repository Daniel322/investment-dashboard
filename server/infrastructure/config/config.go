package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type IConfig struct {
	Name   string
	Logger *log.Logger
}

var logger = log.New(os.Stdout, "Config ", log.LstdFlags)

var Config = IConfig{
	Name:   "Config",
	Logger: logger,
}

func (cfg *IConfig) Bootstrap() error {

	err := godotenv.Load()

	if err != nil {
		Config.Logger.Println("Error loading .env file", err)

		return err
	}

	return nil
}

func (cfg *IConfig) get(field string) (string, error) {
	field_value := os.Getenv(field)

	if len(field_value) == 0 {
		return field_value, errors.New("error on get " + field)
	}

	return field_value, nil
}

func (cfg *IConfig) DB_URL() (string, error) {
	return cfg.get(DB_URL)
}

func (cfg *IConfig) HTTP_PORT() (string, error) {
	return cfg.get(HTTP_PORT)
}

func (cfg *IConfig) BYBIT_API_KEY() (string, error) {
	return cfg.get(BYBIT_API_KEY)
}

func (cfg *IConfig) BYBIT_SECRET_KEY() (string, error) {
	return cfg.get(BYBIT_SECRET_KEY)
}
