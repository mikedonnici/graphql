package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/mikedonnici/graphql/data"
	"github.com/mikedonnici/graphql/schema/types"
)

var Post = &graphql.Field{
	Name:        "Post",
	Description: "A single post",
	Type:        types.Post,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type:        &graphql.NonNull{OfType: graphql.Int},
			Description: "Unique id of the post",
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, ok := p.Args["id"].(int)
		if ok {
			return data.GetPost(id)
		}
		return nil, nil
	},
}
