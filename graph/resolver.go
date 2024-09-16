package graph

import (
	"context"
	"gorm.io/gorm"
	"proyectoIngesoCursos/models"
)

type Resolver struct {
	DB *gorm.DB
}

func (r *Resolver) Cursos(ctx context.Context) ([]*models.Curso, error) {
	var cursos []*models.Curso
	if err := r.DB.Preload("Instructor").Find(&cursos).Error; err != nil {
		return nil, err
	}
	return cursos, nil
}

func (r *Resolver) Curso(ctx context.Context, courseID string) (*models.Curso, error) {
	var curso models.Curso
	if err := r.DB.Preload("Instructor").First(&curso, "course_id = ?", courseID).Error; err != nil {
		return nil, err
	}
	return &curso, nil
}

func (r *Resolver) CreateCurso(ctx context.Context, instructorID string, title string, description string, price float64, category string) (*models.Curso, error) {
	curso := &models.Curso{
		CourseID:     generateID(),
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

// Función auxiliar para generar un ID (puedes usar una librería UUID aquí)
func generateID() string {
	return "some-unique-id" // Implementa tu lógica de generación de IDs aquí
}
