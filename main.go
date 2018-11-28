package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/jmaeso/graphql-test/app/api"
	"github.com/jmaeso/graphql-test/app/graphql"
	"github.com/jmaeso/graphql-test/app/storage/postgres"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:@127.0.0.1/graphql-test_dev?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	ordersStore := &postgres.OrdersStore{
		SQL: db,
	}

	ordersSchema, err := graphql.NewOrdersSchema(ordersStore)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/orders", api.OrdersHandler(ordersSchema))

	http.ListenAndServe(":8080", nil)
}
