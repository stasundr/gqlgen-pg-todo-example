package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"fmt"

	"github.com/stasundr/gqlgen-pg-todo-example/graph/model"
)

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, id int) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todo - todo"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context, limit *int, offset *int) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
