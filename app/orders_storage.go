package app

import "errors"

var ErrOrderNotFound = errors.New("Order Not Found")

type OrdersStore interface {
	Insert(*Order) error
	GetByID(string) (*Order, error)
}
