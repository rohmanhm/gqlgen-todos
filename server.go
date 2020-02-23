package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rohmanhm/gqlgen-todos/graph/generated"
	"github.com/rohmanhm/gqlgen-todos/resolver"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolvers := resolver.NewRoot()
	config := generated.Config{Resolvers: resolvers}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
