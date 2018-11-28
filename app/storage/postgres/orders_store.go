package postgres

import (
	"database/sql"

	"github.com/jmaeso/graphql-test/app"
)

const (
	findOrderByIDStmt = "SELECT id, retailer_id, num_packages FROM orders WHERE id = $1"
	insertOrderStmt   = "INSERT INTO orders (id, retailer_id, num_packages) VALUES ($1, $2, $3)"
)

// OrdersStore must satisfy app.OrdersStore interface.
// Responsible of the storage of app.Order in postgresql.
type OrdersStore struct {
	SQL *sql.DB
}

// Insert inserts an app.Order in the DB.
func (os OrdersStore) Insert(order *app.Order) error {
	_, err := os.SQL.Exec(insertOrderStmt, order.ID, order.RetailerID, order.NumPackages)
	if err != nil {
		return err
	}

	return nil
}

// FindByID queries the DB for an app.Order given its ID.
//
// Will return app.ErrOrderNotFound if the order can not be found.
func (os OrdersStore) FindByID(id string) (*app.Order, error) {
	order := new(app.Order)

	if err := os.SQL.QueryRow(findOrderByIDStmt, id).Scan(&order.ID, &order.RetailerID, &order.NumPackages); err != nil {
		if err == sql.ErrNoRows {
			return nil, app.ErrOrderNotFound
		}

		return nil, err
	}

	return order, nil
}
