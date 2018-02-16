package types

import (
	"github.com/graphql-go/graphql"
)

var Voyage = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Voyage",
	Description: "A complete journey, undertaken by a vessel.",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"vesselId": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"waypoints": &graphql.Field{
			Type: graphql.NewList(Waypoint),
		},
		"positions": &graphql.Field{
			Type: graphql.NewList(Position),
		},
	},
})

var VoyageInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "VoyageInput",
	Description: "Voyage input type",
	Fields: graphql.InputObjectConfigFieldMap{
		"vesselId": &graphql.InputObjectFieldConfig{
			Type: &graphql.NonNull{OfType: graphql.String},
		},
		"name": &graphql.InputObjectFieldConfig{
			Type: &graphql.NonNull{OfType: graphql.String},
		},
		// optional
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"waypoints": &graphql.InputObjectFieldConfig{
			Type: &graphql.List{OfType: VoyageWaypointInput},
		},
	},
})
