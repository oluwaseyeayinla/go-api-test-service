//go:generate go run -mod=mod github.com/99designs/gqlgen
package main

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/zalbiraw/go-api-test-service/helpers"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/comments/graph"
	"log"
	"net/http"
	"os"
)

const defaultPort = "4103"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := helpers.LoadComments()

	if nil != err {
		panic("Unable to load comments-subgraph.")
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graph.GraphQLEndpointHandler(graph.EndpointOptions{EnableDebug: false}))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
