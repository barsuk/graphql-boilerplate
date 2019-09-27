package controllers

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"graphql-boilerplate/models/article"
)

var queryObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			//"get_article":     article.Get(),
			"list_article": article.GetList(),
		},
	})
var mutationObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			//"create_article":     article.Create(),
			"create_article": nil,
			//"delete_article":     article.Delete(),
		},
	})
var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryObject,
		Mutation: mutationObject,
	},
)

// GraphQL Schema
func GraphQL(payload *Payload) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  payload.Query,
		VariableValues: payload.Variables,
	})
	if len(result.Errors) > 0 {
		// log graphQL error and query
		//logs.API(result.Errors, payload.Query)
		fmt.Printf("Ошибки: %v", result.Errors)
	}
	return result
}
