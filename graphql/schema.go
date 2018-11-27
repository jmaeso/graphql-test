package graphql

import "github.com/graphql-go/graphql"

func NewSchema(queryType *graphql.Object, mutationType *graphql.Object) (*graphql.Schema, error) {
	s, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
	if err != nil {
		return nil, err
	}

	return &s, nil
}
