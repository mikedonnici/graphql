package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/mikedonnici/graphql/data"
	"github.com/mikedonnici/graphql/schema/types"
)

var AddVessel = &graphql.Field{
	Name:        "AddVessel",
	Description: "Add a vessel",
	Type:        types.Vessel,
	Args: graphql.FieldConfigArgument{
		"vessel": &graphql.ArgumentConfig{
			Type:        types.VesselInput,
			Description: "A vessel object",
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		vessel, ok := p.Args["vessel"].(map[string]interface{})
		if ok {
			dv := data.Vessel{}
			dv.Unpack(vessel)
			return data.AddVessel(dv), nil
		}
		return nil, nil
	},
}
