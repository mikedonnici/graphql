package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/mikedonnici/graphql/data"
	"github.com/mikedonnici/graphql/schema/types"
)

var Vessel = &graphql.Field{
	Name:        "Vessel",
	Description: "Fetch a single Vessel",
	Type:        types.Vessel,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type:        &graphql.NonNull{OfType: graphql.String},
			Description: "Vessel UUID",
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, ok := p.Args["id"].(string)
		if ok {
			return data.GetVessel(id), nil
		}
		return nil, nil
	},
}

var VesselPositions = &graphql.Field{
	Name:        "VesselPositions",
	Description: "Fetch positions for a Vessel",
	Type:        types.VesselPositions,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type:        &graphql.NonNull{OfType: graphql.String},
			Description: "Vessel UUID",
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, ok := p.Args["id"].(string)
		if ok {
			vp := data.GetVesselPositions(id)
			return vp, nil
		}
		return nil, nil
	},
}
