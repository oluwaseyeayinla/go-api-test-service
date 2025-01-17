//go:generate go run -mod=mod github.com/99designs/gqlgen
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/zalbiraw/go-api-test-service/services/graphql/users/graph"
	"github.com/zalbiraw/go-api-test-service/services/graphql/users/helpers"
)

const defaultPort = "4101"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := helpers.LoadUsers()

	if nil != err {
		panic("Unable to load users.")
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
