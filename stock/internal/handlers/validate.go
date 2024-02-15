package handlers

import "errors"

var (
	ErrIncorrectSKU      = errors.New("incorrect SKU")
	ErrIncorrectQuantity = errors.New("incorrect item quantity")
	ErrIncorrectStockId  = errors.New("incorrect StockId")
)
