package asset_type

import "errors"

type AssetType int

const (
	Crypto AssetType = iota
	Stock
	Cash
	Deposit
	Usd
	Gold
)

var ASSET_TYPES = map[AssetType]string{
	Crypto:  "crypto",
	Stock:   "stock",
	Cash:    "cash",
	Deposit: "deposit",
	Usd:     "usd",
	Gold:    "gold",
}

func (t AssetType) String() string {
	return ASSET_TYPES[t]
}

func New(s string) (AssetType, error) {
	switch s {
	case "crypto":
		return Crypto, nil
	case "stock":
		return Stock, nil
	case "cash":
		return Cash, nil
	case "deposit":
		return Deposit, nil
	case "usd":
		return Usd, nil
	case "gold":
		return Gold, nil
	default:
		return -1, errors.New("invalid asset type")
	}
}
