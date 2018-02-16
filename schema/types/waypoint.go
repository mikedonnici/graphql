package types

import "github.com/graphql-go/graphql"

var Waypoint = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Waypoint",
	Description: "An intended position of a Vessel undertaking a Voyage",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: "Unique identifier (uuid)",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "Arbitrary name for the waypoint - eg 'Point Perpendicular lighthouse'",
		},
		"sequence": &graphql.Field{
			Type:        graphql.Int,
			Description: "Position in numerical sequence of waypoints, ordered by ETA ascending.",
		},
		"eta": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "Estimated arrival time at waypoint, in RFC 3339 format - eg '2006-01-02T15:04:05Z07:00'",
		},
		"ata": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "Actual time of arrival, in RFC 3339 format - eg '2006-01-02T15:04:05Z07:00'",
		},
		"lat": &graphql.Field{
			Type:        graphql.Float,
			Description: "Latitude of waypoint in decimal format - eg -35.099220",
		},
		"lon": &graphql.Field{
			Type:        graphql.Float,
			Description: "Longitude of waypoint in decimal format - eg 150.803005",
		},
	},
})

// WaypointInput type for adding a waypoint alone, hence the need to supply
// the voyage id.
var WaypointInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "WaypointInput",
	Description: "Waypoint input type",
	Fields: graphql.InputObjectConfigFieldMap{
		"voyageId": &graphql.InputObjectFieldConfig{
			Type: &graphql.NonNull{OfType: graphql.String},
		},
		"lat": &graphql.InputObjectFieldConfig{
			Type: &graphql.NonNull{OfType: graphql.Float},
		},
		"lon": &graphql.InputObjectFieldConfig{
			Type: &graphql.NonNull{OfType: graphql.Float},
		},
		// optional
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"eta": &graphql.InputObjectFieldConfig{
			Type: graphql.DateTime,
		},
		"ata": &graphql.InputObjectFieldConfig{
			Type: graphql.DateTime,
		},
	},
})

// VoyageWaypointInput type for adding a waypoint in the context of a Voyage, hence no need to
// supply the voyage id.
var VoyageWaypointInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "VoyageWaypointInput",
	Description: "VoyageWaypoint input type",
	Fields: graphql.InputObjectConfigFieldMap{
		"lat": &graphql.InputObjectFieldConfig{
			Type: &graphql.NonNull{OfType: graphql.Float},
		},
		"lon": &graphql.InputObjectFieldConfig{
			Type: &graphql.NonNull{OfType: graphql.Float},
		},
		// optional
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"number": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"eta": &graphql.InputObjectFieldConfig{
			Type: graphql.DateTime,
		},
		"ata": &graphql.InputObjectFieldConfig{
			Type: graphql.DateTime,
		},
	},
})
