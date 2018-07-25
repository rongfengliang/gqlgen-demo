package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rongfengliang/gqlgen-app/graph"
	"github.com/vektah/gqlgen/handler"
)

func main() {
	http.Handle("/", handler.Playground("Todo", "/query"))
	http.Handle("/query", handler.GraphQL(graph.NewExecutableSchema(&graph.App{})))

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
