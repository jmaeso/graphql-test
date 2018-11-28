package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/jmaeso/graphql-test/app"
	"github.com/pborman/uuid"
)

func NewOrdersSchema(ordersStore app.OrdersStore) (graphql.Schema, error) {
	queryType := newOrdersQueryType(ordersStore)
	mutationType := newOrdersMutationType(ordersStore)

	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
}

var orderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Order",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if order, ok := p.Source.(*app.Order); ok == true {
					return order.ID, nil
				}
				return nil, nil
			},
		},
		"retailer_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if order, ok := p.Source.(*app.Order); ok == true {
					return order.RetailerID, nil
				}
				return nil, nil
			},
		},
		"num_packages": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if order, ok := p.Source.(*app.Order); ok == true {
					return order.NumPackages, nil
				}
				return nil, nil
			},
		},
	},
})

func newOrdersQueryType(ordersStore app.OrdersStore) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"order": &graphql.Field{
				Type: orderType,
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

func newOrdersMutationType(ordersStore app.OrdersStore) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createOrder": &graphql.Field{
				Type: orderType,
				Args: graphql.FieldConfigArgument{
					"retailer_id": &graphql.ArgumentConfig{
						Description: "Retailer's order_id",
						Type:        graphql.NewNonNull(graphql.String),
					},
					"num_packages": &graphql.ArgumentConfig{
						Description: "Order's number of packages",
						Type:        graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					retailerID := p.Args["retailer_id"].(string)
					numPackages := p.Args["num_packages"].(int)

					order := &app.Order{
						ID:          uuid.New(),
						RetailerID:  retailerID,
						NumPackages: numPackages,
					}

					err := ordersStore.Insert(order)
					return order, err
				},
			},
		},
	})
}
