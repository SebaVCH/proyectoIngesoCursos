package graph

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"proyectoIngesoCursos/graph/model"
	"proyectoIngesoCursos/models"
)

type Resolver struct {
	DB *gorm.DB
}

// Obtener todos los cursos
func (r *Resolver) Cursos(ctx context.Context) ([]*models.Curso, error) {
	var cursos []*models.Curso
	if err := r.DB.Preload("Instructor").Find(&cursos).Error; err != nil {
		return nil, err
	}
	return cursos, nil
}

// Obtener un curso por ID
func (r *Resolver) Curso(ctx context.Context, courseID string) (*models.Curso, error) {
	var curso models.Curso
	if err := r.DB.Preload("Instructor").First(&curso, "course_id = ?", courseID).Error; err != nil {
		return nil, err
	}
	return &curso, nil
}

// Crear un nuevo curso
func (r *Resolver) CreateCurso(ctx context.Context, instructorID string, title string, description string, price float64, category string) (*model.Curso, error) {
	curso := &model.Curso{
		InstructorID: instructorID,
		Title:        title,
		Description:  description,
		Price:        price,
		Category:     category,
	}
	if err := r.DB.Create(curso).Error; err != nil {
		return nil, err
	}
	return curso, nil
}

// Actualizar un curso existente
func (r *Resolver) UpdateCurso(ctx context.Context, courseID string, title *string, description *string, price *float64, category *string) (*models.Curso, error) {
	var curso models.Curso
	if err := r.DB.First(&curso, "course_id = ?", courseID).Error; err != nil {
		return nil, err
	}

	// Actualizar los campos si no son nulos
	if title != nil {
		curso.Title = *title
	}
	if description != nil {
		curso.Description = *description
	}
	if price != nil {
		curso.Price = *price
	}
	if category != nil {
		curso.Category = *category
	}

	if err := r.DB.Save(&curso).Error; err != nil {
		return nil, err
	}
	return &curso, nil
}

// Eliminar un curso
func (r *Resolver) DeleteCurso(ctx context.Context, courseID string) (bool, error) {
	var curso models.Curso
	if err := r.DB.First(&curso, "course_id = ?", courseID).Error; err != nil {
		return false, err
	}

	if err := r.DB.Delete(&curso).Error; err != nil {
		return false, err
	}
	return true, nil
}

// Buscar cursos por título o categoría
func (r *Resolver) SearchCursos(ctx context.Context, query string) ([]*models.Curso, error) {
	var cursos []*models.Curso
	if err := r.DB.Preload("Instructor").
		Where("title ILIKE ? OR category ILIKE ?", "%"+query+"%", "%"+query+"%").
		Find(&cursos).Error; err != nil {
		return nil, err
	}
	return cursos, nil
}

func (r *queryResolver) CourseByID(ctx context.Context, courseID int) (*model.Curso, error) {
	var course model.Curso
	// Buscar el curso por ID en la base de datos
	if err := r.DB.Where("course_id = ?", courseID).First(&course).Error; err != nil {
		return nil, fmt.Errorf("curso no encontrado")
	}

	return &course, nil
}

// Función auxiliar para generar un ID único
func generateUniqueID() string {
	return uuid.NewString()
}
