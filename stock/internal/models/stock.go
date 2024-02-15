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

func (s Status) IsAvailable() bool {
	switch s {
	case StatusAvailable:
		return true
	case StatusNotAvailable:
		return false
	default:
		return false
	}
}

type ItemStock struct {
	SKU       uint32
	StockId   uint32
	Available uint32
	Reserved  uint32
}
