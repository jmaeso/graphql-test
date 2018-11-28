package postgres

import (
	"database/sql"

	"github.com/jmaeso/graphql-test/app"
)

const (
	findOrderByIDStmt = "SELECT id, retailer_id, num_packages FROM orders WHERE id = $1"
	insertOrderStmt   = "INSERT INTO orders (id, retailer_id, num_packages) VALUES ($1, $2, $3)"
)

type OrdersStore struct {
	SQL *sql.DB
}

func (os OrdersStore) Insert(order *app.Order) error {
	_, err := os.SQL.Exec(insertOrderStmt, order.ID, order.RetailerID, order.NumPackages)
	if err != nil {
		return err
	}

	return nil
}

func (os OrdersStore) GetByID(id string) (*app.Order, error) {
	order := new(app.Order)

	if err := os.SQL.QueryRow(findOrderByIDStmt, id).Scan(&order.ID, &order.RetailerID, &order.NumPackages); err != nil {
		if err == sql.ErrNoRows {
			return nil, app.ErrOrderNotFound
		}

		return nil, err
	}

	return order, nil
}
