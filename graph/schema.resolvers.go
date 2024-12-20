package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"
	"fmt"
	"proyectoIngesoCursos/graph/model"
	"proyectoIngesoCursos/models"
	"strconv"
)

// CreateCurso is the resolver for the createCurso field.
func (r *mutationResolver) CreateCurso(ctx context.Context, title string, description string, price int, category string, imageURL string, instructorName string) (*model.Curso, error) {
	return r.Resolver.CreateCurso(ctx, title, description, price, category, imageURL, instructorName)
}

// DeleteCursoByID is the resolver for the deleteCursoByID field.
func (r *mutationResolver) DeleteCursoByID(ctx context.Context, courseID int) (string, error) {
	return r.Resolver.DeleteCursoByID(ctx, courseID)
}

// UpdateCursoByID is the resolver for the updateCursoByID field.
func (r *mutationResolver) UpdateCursoByID(ctx context.Context, courseID int, title string, description string, price int, category string, imageURL string) (*model.Curso, error) {
	updatedCurso, err := r.Resolver.UpdateCursoByID(ctx, courseID, title, description, price, category, imageURL)
	if err != nil {
		return nil, err
	}

	// Convertir el modelo Gorm a GraphQL
	return &model.Curso{
		CourseID:       int(updatedCurso.CourseID),
		InstructorName: updatedCurso.InstructorName,
		Title:          updatedCurso.Title,
		Description:    updatedCurso.Description,
		Price:          updatedCurso.Price,
		Category:       updatedCurso.Category,
		ImageURL:       updatedCurso.ImageURL,
	}, nil
}

// Cursos is the resolver for the cursos field.
func (r *queryResolver) Cursos(ctx context.Context) ([]*model.Curso, error) {
	var cursos []models.Curso

	// Obtener todos los cursos de la base de datos
	if err := r.DB.Find(&cursos).Error; err != nil {
		return nil, fmt.Errorf("error al obtener los cursos: %v", err)
	}

	// Convertir los modelos de la base de datos al modelo GraphQL
	var result []*model.Curso
	for _, curso := range cursos {
		result = append(result, &model.Curso{
			CourseID:       int(curso.CourseID),
			InstructorName: curso.InstructorName,
			Title:          curso.Title,
			Description:    curso.Description,
			Price:          curso.Price,
			Category:       curso.Category,
			ImageURL:       curso.ImageURL,
		})
	}

	return result, nil
}

// Curso is the resolver for the curso field.
func (r *queryResolver) Curso(ctx context.Context, courseID string) (*model.Curso, error) {
	return r.CursoByID(ctx, courseID)
}

// CursoByID is the resolver for the cursoByID field.
func (r *queryResolver) CursoByID(ctx context.Context, courseID string) (*model.Curso, error) {
	// Convertir el courseID de string a uint
	id, err := strconv.ParseUint(courseID, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("ID inválido: %v", err)
	}

	var curso models.Curso

	// Buscar el curso por ID en la base de datos
	if err := r.DB.First(&curso, "course_id = ?", uint(id)).Error; err != nil {
		return nil, fmt.Errorf("curso no encontrado")
	}

	// Convertir el modelo de la base de datos al modelo GraphQL
	return &model.Curso{
		CourseID:       int(curso.CourseID),
		InstructorName: curso.InstructorName,
		Title:          curso.Title,
		Description:    curso.Description,
		Price:          curso.Price,
		Category:       curso.Category,
		ImageURL:       curso.ImageURL,
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
