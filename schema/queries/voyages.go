package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/mikedonnici/graphql/data"
	"github.com/mikedonnici/graphql/schema/types"
)

var Voyages = &graphql.Field{
	Type: graphql.NewList(types.Voyage),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return data.GetVoyages(), nil
	},
}
