package models

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        int    `gorm:"primaryKey"`
	FirstName string `gorm:"size:100"`
	LastName  string `gorm:"size:100"`
	Email     string `gorm:"size:100;uniqueIndex"`
	Password  string `gorm:"size:255"`
}

