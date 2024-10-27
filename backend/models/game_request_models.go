package models

import "time"

type GameRequest struct {
	ID        int    `gorm:"primaryKey"`
	UserEmail string `gorm:"size:100"`
	Niveau    int
	Location  string `gorm:"size:100"`
	Time      time.Time
	Gender    string `gorm:"size:50"`
	Amount    int
	Price     float64 `gorm:"type:decimal(10,2)"`
}
