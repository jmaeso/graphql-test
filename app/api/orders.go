package api

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func OrdersHandler(ordersSchema graphql.Schema) http.Handler {
	return handler.New(&handler.Config{
		Schema: &ordersSchema,
		Pretty: true,
		// GraphiQL: true,
	})
}
