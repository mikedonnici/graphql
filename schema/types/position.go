package types

import (
	"github.com/graphql-go/graphql"
)

var Position = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Position",
	Description: "A reported position of a vessel at a given time",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"vesselId": &graphql.Field{
			Type: graphql.String,
		},
		"ata": &graphql.Field{
			Type: graphql.DateTime,
		},
		"sequence": &graphql.Field{
			Type: graphql.Int,
		},
		"lat": &graphql.Field{
			Type: graphql.Float,
		},
		"lon": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

var PositionInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "PositionInput",
	Description: "Position input type",
	Fields: graphql.InputObjectConfigFieldMap{
		"vesselId": &graphql.InputObjectFieldConfig{
			Type: &graphql.NonNull{OfType: graphql.String},
		},
		"ata": &graphql.InputObjectFieldConfig{
			Type: &graphql.NonNull{OfType: graphql.DateTime},
		},
		"lat": &graphql.InputObjectFieldConfig{
			Type: &graphql.NonNull{OfType: graphql.Float},
		},
		"lon": &graphql.InputObjectFieldConfig{
			Type: &graphql.NonNull{OfType: graphql.Float},
		},
	},
})
