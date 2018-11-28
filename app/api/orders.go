package api

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// OrdersHandler is the responsible to handle all the requests to the /orders URL.
func OrdersHandler(ordersSchema graphql.Schema) http.Handler {
	return handler.New(&handler.Config{
		Schema: &ordersSchema,
		Pretty: true,
		// GraphiQL: true,
	})
}
