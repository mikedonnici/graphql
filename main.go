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
				"posts":   queries.Posts,
				"post":    queries.Post,
				"authors": queries.Authors,
				"author":  queries.Author,
				"authorPosts": queries.AuthorPosts,
			},
		})

	// example mutation
	// https://gist.github.com/sogko/7debd336118e5e7c7f65

	rootMutation := graphql.NewObject(
		graphql.ObjectConfig{
			Name:        "RootMutation",
			Description: "...",
			Fields: graphql.Fields{
				"addPost": mutations.AddPost,
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
