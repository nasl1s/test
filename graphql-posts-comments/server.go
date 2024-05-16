// server.go

package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"

	"github.com/nasl1s/test/graphql-posts-comments/db"
	"github.com/nasl1s/test/graphql-posts-comments/generated"
	"github.com/nasl1s/test/graphql-posts-comments/graph"
)

const defaultPort = "8080"

func main() {
	db := db.NewInMemoryDB()

	resolver := &graph.Resolver{
		db: db,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.MultipartForm{})

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	srv.AddTransport(websocket.New(upgrader))

	srv.SetQueryCache(lru.New(1000))

	http.Handle("/", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
