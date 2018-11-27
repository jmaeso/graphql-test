package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/jmaeso/graphql-test/app"
)

var OrderType = graphql.NewObject(graphql.ObjectConfig{
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
