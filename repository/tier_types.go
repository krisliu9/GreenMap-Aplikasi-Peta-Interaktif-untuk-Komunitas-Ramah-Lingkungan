package repository

import "time"

type Tier struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Tier_Name     string    `json:"tier_name" gorm:"column:tier_name"`
	Minimal_Point int       `json:"minimal_point" gorm:"column:minimal_point"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"column:updated_at"`
}
