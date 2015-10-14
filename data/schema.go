package data

import (
	"github.com/chris-ramon/graphql-go/types"
	"github.com/sogko/graphql-relay-go"
)

// In GraphQL's type system notation, here is our schema
// type Post {
//   id:   ID!
//   text: String
// }
// type Query {
//	 latestPost: Post
// }

var postType *types.GraphQLObjectType
var queryType *types.GraphQLObjectType
var Schema types.GraphQLSchema

func init() {
	postType = types.NewGraphQLObjectType(types.GraphQLObjectTypeConfig{
		Name: "Post",
		Fields: types.GraphQLFieldConfigMap{
			// Define `id` field as a Relay GlobalID field.
			// It helps with translating your GraphQL object's id into a global id
			// For eg:
			// 		For a `Post` type, with an id of `1`, it's global id will be `UG9zdDox`
			//      which is a base64 encoded version of `Post:1` string
			// We will explore more in the next part of this series.
			"id": gqlrelay.GlobalIdField("Post", nil),
			"text": &types.GraphQLFieldConfig{
				Type: types.GraphQLString,
			},
		},
	})

	queryType = types.NewGraphQLObjectType(types.GraphQLObjectTypeConfig{
		Name: "Query",
		Fields: types.GraphQLFieldConfigMap{
			"latestPost": &types.GraphQLFieldConfig{
				Type: postType,
				Resolve: func(p types.GQLFRParams) interface{} {
					return GetLatestPost()
				},
			},
		},
	})

	Schema, _ = types.NewGraphQLSchema(types.GraphQLSchemaConfig{
		Query: queryType,
	})

}