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
<<<<<<< HEAD
func (r *Resolver) CreateCurso(ctx context.Context, title, description string, price float64, category, imageURL string, instructorName string) (*model.Curso, error) {
	// Crear el modelo del curso para la base de datos
	cursoDB := models.Curso{
		InstructorName: instructorName,
		Title:          title,
		Description:    description,
		Price:          price,
		Category:       category,
		ImageURL:       imageURL,
=======
func (r *Resolver) CreateCurso(ctx context.Context, instructorID, title, description string, price float64, category string) (*model.Curso, error) {
	// Crear el modelo del curso para la base de datos
	cursoDB := models.Curso{
		InstructorID: instructorID,
		Title:        title,
		Description:  description,
		Price:        price,
		Category:     category,
>>>>>>> parent of cd0bb11 (Logica para mostrar imagen de cada curso y mostrar el nombre del instructor)
	}

	// Guardar el curso en la base de datos
	if err := r.DB.Create(&cursoDB).Error; err != nil {
		return nil, err
	}

	// Convertir el modelo de la base de datos al modelo GraphQL
	cursoGraphQL := &model.Curso{
<<<<<<< HEAD
		CourseID:       int(cursoDB.CourseID),
		InstructorName: cursoDB.InstructorName,
		Title:          cursoDB.Title,
		Description:    cursoDB.Description,
		Price:          cursoDB.Price,
		Category:       cursoDB.Category,
		ImageURL:       cursoDB.ImageURL,
=======
		CourseID:     int(cursoDB.CourseID), // Conversión de uint a int
		InstructorID: cursoDB.InstructorID,
		Title:        cursoDB.Title,
		Description:  cursoDB.Description,
		Price:        cursoDB.Price,
		Category:     cursoDB.Category,
>>>>>>> parent of cd0bb11 (Logica para mostrar imagen de cada curso y mostrar el nombre del instructor)
	}

	return cursoGraphQL, nil
}

// UpdateCursoByID - actualiza las variables de un curso por su ID
func (r *Resolver) UpdateCursoByID(ctx context.Context, courseID uint, title, description string, price float64, category string) (*models.Curso, error) {
	var curso models.Curso

	// Buscar el curso por ID
	if err := r.DB.First(&curso, "course_id = ?", courseID).Error; err != nil {
		return nil, fmt.Errorf("curso no encontrado")
	}

	// Actualizar las variables del curso
	curso.Title = title
	curso.Description = description
	curso.Price = price
	curso.Category = category

	// Guardar los cambios
	if err := r.DB.Save(&curso).Error; err != nil {
		return nil, fmt.Errorf("no se pudo actualizar el curso")
	}

	return &curso, nil
}

// DeleteCursoByID - elimina un curso por su ID
func (r *Resolver) DeleteCursoByID(ctx context.Context, courseID uint) (string, error) {
	var curso models.Curso

	// Buscar el curso por ID
	if err := r.DB.First(&curso, "course_id = ?", courseID).Error; err != nil {
		return "", fmt.Errorf("curso no encontrado")
	}

	// Eliminar el curso de la base de datos
	if err := r.DB.Delete(&curso).Error; err != nil {
		return "", fmt.Errorf("no se pudo eliminar el curso")
	}

	return "Curso eliminado exitosamente", nil
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
