package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/jmaeso/graphql-test/app"
)

func NewQueryType(ordersStore app.OrdersStore) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"order": &graphql.Field{
				Type: OrderType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "Order ID",
						Type:        graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					i := p.Args["id"].(string)
					return ordersStore.GetByID(i)
				},
			},
		},
	})
}
