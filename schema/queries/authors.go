package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/mikedonnici/graphql/data"
	"github.com/mikedonnici/graphql/schema/types"
)

var Authors = &graphql.Field{
	Type: graphql.NewList(types.Author),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return data.GetAuthors(), nil
	},
}
