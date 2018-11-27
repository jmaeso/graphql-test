package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/jmaeso/graphql-test/graphql"
	"github.com/jmaeso/graphql-test/storage/postgres"
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

	queryType := graphql.NewQueryType(ordersStore)
	mutationType := graphql.NewMutationType(ordersStore)

	schema, err := graphql.NewSchema(queryType, mutationType)
	if err != nil {
		log.Fatal(err)
	}

	// create a graphl-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := handler.New(&handler.Config{
		Schema:   schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)

	// and serve!
	http.ListenAndServe(":8080", nil)
}
