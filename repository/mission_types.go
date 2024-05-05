package repository

import (
	"time"
)

type Mission struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Target      int       `json:"target" gorm:"column:target"`
	Description string    `json:"description" gorm:"column:description"`
	Point       int       `json:"point" gorm:"column:point"`
	StartAt     time.Time `json:"start_at" gorm:"column:start_at"`
	EndAt       time.Time `json:"end_at" gorm:"column:end_at"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}
