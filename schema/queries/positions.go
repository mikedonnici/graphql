package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/mikedonnici/graphql/data"
	"github.com/mikedonnici/graphql/schema/types"
)

var Positions = &graphql.Field{
	Type: graphql.NewList(types.Position),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return data.GetPositions(), nil
	},
}
