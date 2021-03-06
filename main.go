package main

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/mikedonnici/graphql/schema/mutations"
	"github.com/mikedonnici/graphql/schema/queries"
)

func main() {

	rootQuery := graphql.NewObject(
		graphql.ObjectConfig{
			Name:        "RootQuery",
			Description: "...",
			Fields: graphql.Fields{
				"vessels":         queries.Vessels,
				"vessel":          queries.Vessel,
				"vesselPositions": queries.VesselPositions,
				"voyages":         queries.Voyages,
				//"voyage": queries.Voyage,
				"positions": queries.Positions,
			},
		})

	rootMutation := graphql.NewObject(
		graphql.ObjectConfig{
			Name:        "RootMutation",
			Description: "...",
			Fields: graphql.Fields{
				"addVessel":   mutations.AddVessel,
				"addVoyage":   mutations.AddVoyage,
				"addPosition": mutations.AddPosition,
			},
		})

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    rootQuery,
			Mutation: rootMutation,
		},
	)
	if err != nil {
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/graphql", h)
	http.ListenAndServe(":8088", nil)

}
