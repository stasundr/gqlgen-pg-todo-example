package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"fmt"
	"time"

	"github.com/stasundr/gqlgen-pg-todo-example/graph/model"
)

// TodoCreate is the resolver for the todoCreate field.
func (r *mutationResolver) TodoCreate(ctx context.Context, todo model.TodoInput) (*model.Todo, error) {
	// Validate that createdby id actually exists
	user := model.User{ID: todo.CreatedBy}
	if err := r.DB.NewSelect().Model(&user).Where("id = ?", todo.CreatedBy).Scan(ctx); err != nil {
		return nil, err
	}

	t := model.Todo{
		Name:      todo.Name,
		CreatedBy: todo.CreatedBy,
		UpdatedBy: todo.CreatedBy,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := r.DB.NewInsert().Model(&t).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// TodoComplete is the resolver for the todoComplete field.
func (r *mutationResolver) TodoComplete(ctx context.Context, id int, updatedBy int) (*model.Todo, error) {
	// Validate that updatedBy id actually exists
	user := model.User{ID: updatedBy}
	if err := r.DB.NewSelect().Model(&user).Where("id = ?", updatedBy).Scan(ctx); err != nil {
		return nil, fmt.Errorf("user %d does not exist", updatedBy)
	}

	todo := model.Todo{ID: id}
	if err := r.DB.NewSelect().Model(&todo).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, fmt.Errorf("todo %d does not exist", id)
	}

	todo.UpdatedBy = updatedBy
	todo.IsComplete = true
	todo.UpdatedAt = time.Now()

	_, err := r.DB.NewUpdate().Model(&todo).WherePK().Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

// TodoDelete is the resolver for the todoDelete field.
func (r *mutationResolver) TodoDelete(ctx context.Context, id int, updatedBy int) (*model.Todo, error) {
	// Validate that updatedBy id actually exists
	user := model.User{ID: updatedBy}
	if err := r.DB.NewSelect().Model(&user).Where("id = ?", updatedBy).Scan(ctx); err != nil {
		return nil, fmt.Errorf("user %d does not exist", updatedBy)
	}

	todo := model.Todo{ID: id}
	if err := r.DB.NewSelect().Model(&todo).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, fmt.Errorf("todo %d does not exist", id)
	}

	todo.UpdatedBy = updatedBy
	todo.IsDeleted = true
	todo.UpdatedAt = time.Now()

	_, err := r.DB.NewUpdate().Model(&todo).WherePK().Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

// UserCreate is the resolver for the userCreate field.
func (r *mutationResolver) UserCreate(ctx context.Context, user model.UserInput) (*model.User, error) {
	usr := model.User{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err := r.DB.NewInsert().Model(&usr).Exec(ctx); err != nil {
		return nil, err
	}

	return &usr, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
