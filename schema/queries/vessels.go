package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/mikedonnici/graphql/data"
	"github.com/mikedonnici/graphql/schema/types"
)

var Vessels = &graphql.Field{
	Type: graphql.NewList(types.Vessel),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return data.GetVessels(), nil
	},
}
