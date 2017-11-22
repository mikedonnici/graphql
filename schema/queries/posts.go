package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/mikedonnici/graphql/data"
	"github.com/mikedonnici/graphql/schema/types"
)

var Posts = &graphql.Field{
	Type: graphql.NewList(types.Post),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return data.GetPosts(), nil
	},
}
