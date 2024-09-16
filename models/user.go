package models

type Usuario struct {
	UserID   string `gorm:"primaryKey;type:text" json:"userID"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
