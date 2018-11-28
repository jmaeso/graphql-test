package app

import "errors"

// ErrOrderNotFound is a typed error to use when a db does not find an Order.
var ErrOrderNotFound = errors.New("Order Not Found")

// OrdersStore defines all the methods that any storage system needs to
// define in order to operate with app.Order.
type OrdersStore interface {
	Insert(*Order) error
	GetByID(string) (*Order, error)
}
