//go:generate go run -mod=mod github.com/99designs/gqlgen
package main

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/zalbiraw/go-api-test-service/helpers"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/users/graph"
	"log"
	"net/http"
	"os"
)

const defaultPort = "4101"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := helpers.LoadUsers()

	if nil != err {
		panic("Unable to load users-subgraph.")
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graph.GraphQLEndpointHandler(graph.EndpointOptions{EnableDebug: false}))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
