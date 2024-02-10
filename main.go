package main

import (
	"fmt"
	"net/http"

	"database/sql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fatih/color"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"github.com/stasundr/gqlgen-pg-todo-example/dataloaders"
	database "github.com/stasundr/gqlgen-pg-todo-example/db"
	"github.com/stasundr/gqlgen-pg-todo-example/graph"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

const (
	port = ":8080"
)

func lineSeparator() {
	fmt.Println("========")
}

func startMessage() {
	lineSeparator()
	color.Green("Listening on localhost%s\n", port)
	color.Green("Visit `http://localhost%s/graphql` in your browser\n", port)
	lineSeparator()
}

func main() {
	lineSeparator()

	// Connect to the database using Bun and the PostgreSQL driver.
	sqldb, err := sql.Open("postgres", "postgres://postgres@localhost:5432/todos?sslmode=disable")
	if err != nil {
		panic(err)
	}

	// Create a Bun DB instance using the PostgreSQL dialect.
	db := bun.NewDB(sqldb, pgdialect.New())

	// Optionally, add a logger to Bun for debugging purposes.
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	defer db.Close()

	err = database.Seed(db)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	// The base path that users would use is POST /graphql which is fairly
	// idiomatic.
	r.Route("/graphql", func(r chi.Router) {
		// Initialize the dataloaders as middleware into our route
		r.Use(dataloaders.NewMiddleware(db)...)

		schema := graph.NewExecutableSchema(graph.Config{
			Resolvers: &graph.Resolver{
				DB: db,
			},
			Directives: graph.DirectiveRoot{},
			Complexity: graph.ComplexityRoot{},
		})

		srv := handler.NewDefaultServer(schema)
		srv.Use(extension.FixedComplexityLimit(300))

		r.Handle("/", srv)
	})

	gqlPlayground := playground.Handler("GraphQL Playground", "/graphql")
	r.Get("/", gqlPlayground)

	startMessage()
	panic(http.ListenAndServe(port, r))
}
