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

func corsMiddleware(next http.Handler) http.Handler {
	allowedOrigins := map[string]bool{
		"http://localhost:3000":                                             true,
		"https://mindbox-frontend.vercel.app":                               true,
		"https://mindbox-frontend-otwblm0v1-j-v-s-cs-projects.vercel.app": true,
		}


	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	db := database.Connect()
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		panic("Application PORT is required")
	}

	roadmapDB := database.NewRoadmapRepository(db)
	categoryDB := database.NewCategoryRepository(db)
	taskDB := database.NewTaskRepository(db)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		RoadmapDB:  roadmapDB,
		CategoryDB: categoryDB,
		TaskDB:     taskDB,
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

	log.Fatal(http.ListenAndServe(":"+port, corsMiddleware(http.DefaultServeMux)))
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

