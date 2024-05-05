package repository

import "time"

type Pinpoint struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"column:name"`
	Description string    `json:"description" gorm:"column:description"`
	Latitude    float64   `json:"latitude" gorm:"column:latitude"`
	Longitude   float64   `json:"longitude" gorm:"column:longitude"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}
