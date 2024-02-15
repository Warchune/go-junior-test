package models

type Stock struct {
	Id     uint32
	Name   string
	Status Status
}

type Status string

const (
	StatusAvailable    Status = "available"
	StatusNotAvailable Status = "not available"
)

type ItemStock struct {
	Sku       uint32
	StockId   uint32
	Available uint32
	Reserved  uint32
}
