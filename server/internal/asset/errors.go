package asset

import "errors"

var (
	ErrAssetPointerIsNil = errors.New("Asset pointer is nil")
	ErrAssetNotFound     = errors.New("Asset not found")
)
