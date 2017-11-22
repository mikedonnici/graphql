package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/mikedonnici/graphql/data"
	"github.com/mikedonnici/graphql/schema/types"
)

var AddPost = &graphql.Field{
	Name:        "AddPost",
	Description: "Add a post",
	Type:        types.Post,
	Args: graphql.FieldConfigArgument{
		"post": &graphql.ArgumentConfig{
			Type:        &graphql.NonNull{OfType: types.Post},
			Description: "The post data",
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		post, ok := p.Args["post"].(data.Post)
		if ok {
			return data.AddPost(post), nil
		}
		return nil, nil
	},
}
