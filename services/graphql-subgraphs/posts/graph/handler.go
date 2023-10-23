package graph

import (
	"net/http"

	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/posts/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
)

type EndpointOptions struct {
	EnableDebug bool
}

var TestOptions = EndpointOptions{
	EnableDebug: false,
}

func GraphQLEndpointHandler(opts EndpointOptions) http.Handler {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{}}))
	if opts.EnableDebug {
		srv.Use(&debug.Tracer{})
	}

	return srv
}
