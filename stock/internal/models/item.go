package models

type Item struct {
	Name         string
	Size         string
	SKU          uint32
	AvailableAll uint32
	ReservedAll  uint32
	StocksId     []uint32
}

type ItemActions struct {
	Name    string
	Size    string
	SKU     uint32
	Count   uint32
	StockId uint32
}

type ItemArrival ItemActions
type ItemReserve ItemActions
type ItemReserveCancel ItemActions
