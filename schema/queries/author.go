package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/mikedonnici/graphql/data"
	"github.com/mikedonnici/graphql/schema/types"
)

var Author = &graphql.Field{
	Type:        types.Author,
	Description: "A single author",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type:        &graphql.NonNull{OfType: graphql.Int},
			Description: "Unique id of the author",
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, ok := p.Args["id"].(int)
		if ok {
			return data.GetAuthor(id), nil
		}
		return nil, nil
	},
}

var AuthorPosts = &graphql.Field{
	Type:        &graphql.List{OfType: types.Post},
	Description: "A list of Post by a particular Author",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type:        &graphql.NonNull{OfType: graphql.Int},
			Description: "id of the author",
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, ok := p.Args["id"].(int)
		if ok {
			return data.GetAuthorPosts(id), nil
		}
		return nil, nil
	},
}
