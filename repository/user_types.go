package repository

import "time"

type User struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name" gorm:"column:name"`
	Email         string    `json:"email" gorm:"column:email"`
	Current_Point int       `json:"current_point" gorm:"column:current_point"`
	Tier_ID       uint      `json:"tier" gorm:"column:tier_id"`
	Tier_Name     string    `json:"tier_name"`
	Password      string    `json:"password" gorm:"column:password"`
	Role          string    `json:"role" gorm:"column:role"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"column:updated_at"`
}
