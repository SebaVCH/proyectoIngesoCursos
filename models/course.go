package models

type Curso struct {
	CourseID       uint    `gorm:"primaryKey;autoIncrement" json:"courseID"` // Cambiado a `uint` y agregado `autoIncrement`
	InstructorID   string  `gorm:"type:text" json:"instructorID"`
	Title          string  `gorm:"type:text" json:"title"`
	Description    string  `gorm:"type:text" json:"description"`
	Price          float64 `gorm:"type:float" json:"price"`
	Category       string  `gorm:"type:text" json:"category"`
	ImageURL       string  `gorm:"type:text" json:"imageURL"`
	InstructorName string  `gorm:"type:text" json:"instructorName"`
}
