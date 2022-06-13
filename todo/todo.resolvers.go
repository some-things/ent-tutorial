package todo

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"todo/ent"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, todo TodoInput) (*ent.Todo, error) {
	return r.client.Todo.Create().
		SetText(todo.Text).
		SetStatus(todo.Status).
		SetNillablePriority(todo.Priority). // Set the "priority" field if provided.
		SetNillableParentID(todo.Parent).   // Set the "parent_id" field if provided.
		Save(ctx)
}

func (r *queryResolver) Todos(ctx context.Context) ([]*ent.Todo, error) {
	return r.client.Todo.Query().All(ctx)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
