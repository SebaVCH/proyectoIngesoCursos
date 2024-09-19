package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.53

import (
	"context"
	"fmt"
	"proyectoIngesoCursos/graph/model"
)

// CreateCurso is the resolver for the createCurso field.
func (r *mutationResolver) CreateCurso(ctx context.Context, instructorID string, title string, description string, price float64, category string) (*model.Curso, error) {
	panic(fmt.Errorf("not implemented: CreateCurso - createCurso"))
}

// Cursos is the resolver for the cursos field.
func (r *queryResolver) Cursos(ctx context.Context) ([]*model.Curso, error) {
	panic(fmt.Errorf("not implemented: Cursos - cursos"))
}

// Curso is the resolver for the curso field.
func (r *queryResolver) Curso(ctx context.Context, courseID string) (*model.Curso, error) {
	panic(fmt.Errorf("not implemented: Curso - curso"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
