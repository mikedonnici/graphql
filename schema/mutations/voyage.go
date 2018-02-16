package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/mikedonnici/graphql/data"
	"github.com/mikedonnici/graphql/schema/types"
)

var AddVoyage = &graphql.Field{
	Name:        "AddVoyage",
	Description: "Add a voyage",
	Type:        types.Voyage,
	Args: graphql.FieldConfigArgument{
		"voyage": &graphql.ArgumentConfig{
			Type:        types.VoyageInput,
			Description: "A voyage object",
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		voyage, ok := p.Args["voyage"].(map[string]interface{})
		if ok {
			dv := data.Voyage{}
			dv.Unpack(voyage)
			// unpack the waypoints, if provided - they are a []interface{}
			waypoints, ok := voyage["waypoints"].([]interface{})
			if ok {
				for _, v := range waypoints {
					waypoint := v.(map[string]interface{})
					dw := data.Waypoint{}
					dw.Unpack(waypoint)
					dv.Waypoints = append(dv.Waypoints, dw)
				}
			}
			dv.Sequence()
			return data.AddVoyage(dv), nil
		}
		return nil, nil
	},
}
