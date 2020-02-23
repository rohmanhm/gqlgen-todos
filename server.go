package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	_ "github.com/rohmanhm/gqlgen-todos/firebase"
	"github.com/rohmanhm/gqlgen-todos/graph/generated"
	"github.com/rohmanhm/gqlgen-todos/resolver"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("could not read the environment: %v", err))
	}

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
