package repository

import (
	"time"
)

type UserMission struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	UserID          uint      `json:"user_id" gorm:"column:user_id"`
	MissionID       uint      `json:"mission_id" gorm:"column:mission_id"`
	CurrentProgress int       `json:"current_progress" gorm:"column:current_progress"`
	CreatedAt       time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"column:updated_at"`
}
