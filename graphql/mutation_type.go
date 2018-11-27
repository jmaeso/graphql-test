package graphql

import (
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/jmaeso/graphql-test/app"
	"github.com/pborman/uuid"
)

func NewMutationType(ordersStore app.OrdersStore) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createOrder": &graphql.Field{
				Type: OrderType,
				Args: graphql.FieldConfigArgument{
					"retailer_id": &graphql.ArgumentConfig{
						Description: "Retailer's order_id",
						Type:        graphql.NewNonNull(graphql.String),
					},
					"num_packages": &graphql.ArgumentConfig{
						Description: "Order's number of packages",
						Type:        graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					retailerID := p.Args["retailer_id"].(string)
					numPackagesStr := p.Args["num_packages"].(string)
					numPackages, err := strconv.Atoi(numPackagesStr)
					if err != nil {
						return nil, err
					}

					order := &app.Order{
						ID:          uuid.New(),
						RetailerID:  retailerID,
						NumPackages: numPackages,
					}

					err = ordersStore.Insert(order)
					return order, err
				},
			},
		},
	})
}
