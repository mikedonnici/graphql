package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/mikedonnici/graphql/data"
	"github.com/mikedonnici/graphql/schema/types"
)

// AddPosition records a new position for a vessel. There is no direct association with
// a Voyage or a Waypoint - it is simply the vessels reported position at a given time
var AddPosition = &graphql.Field{
	Name:        "AddPosition",
	Description: "Add a position for a vessel",
	Type:        types.Position,
	Args: graphql.FieldConfigArgument{
		"position": &graphql.ArgumentConfig{
			Type:        types.PositionInput,
			Description: "A position object",
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		position, ok := p.Args["position"].(map[string]interface{})
		if ok {
			dp := data.Position{}
			dp.Unpack(position)
			return data.AddPosition(dp), nil
		}
		return nil, nil
	},
}
