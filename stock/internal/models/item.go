package models

type Item struct {
	Name      string
	Size      string
	Sku       uint32
	Stocks    []string
	Available uint32
	Reserved  uint32
}

type ItemActions struct {
	Sku   uint32
	Count uint32
	Stock uint32
}
type ItemReserve ItemActions
type ItemReserveCancel ItemActions
