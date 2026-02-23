package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/J-V-S-C/MindBox/graph"
	"github.com/J-V-S-C/MindBox/internal/database"
	_ "github.com/lib/pq"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {

	db := database.Connect()
	defer db.Close()

	port := os.Getenv("PORT")
	if port == ""{
		panic("Application PORT is required")
	}

	roadmapDB := database.NewRoadmap(db)
	categoryDB := database.NewCategory(db)
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		RoadmapDB: roadmapDB,
		CategoryDB: categoryDB,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("you should connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func CheckError(err error){
	if err != nil{
		panic(err)
	}
}