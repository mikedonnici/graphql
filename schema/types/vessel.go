package types

import (
	"github.com/graphql-go/graphql"
)

var Vessel = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Vessel",
	Description: "A ship, boat or other floating entity capable of undertaking a voyage",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"callSign": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var VesselInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "VesselInput",
	Description: "Add new vessel",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: &graphql.NonNull{OfType: graphql.String},
		},
		"callSign": &graphql.InputObjectFieldConfig{
			Type: &graphql.NonNull{OfType: graphql.String},
		},
		"type": &graphql.InputObjectFieldConfig{
			Type: &graphql.NonNull{OfType: graphql.String},
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: &graphql.NonNull{OfType: graphql.String},
		},
	},
})

var VesselPositions = graphql.NewObject(graphql.ObjectConfig{
	Name:        "VesselPositions",
	Description: "Shows positions of a vessel",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"callSign": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"positions": &graphql.Field{
			Type: graphql.NewList(Position),
		},
	},
})
