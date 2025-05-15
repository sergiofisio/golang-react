package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	Admin Role = "admin"
	User  Role = "user"
)

type Users struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"unique;not null"`
	Email     string         `json:"email" gorm:"unique;not null"`
	Password  string         `json:"password" gorm:"not null"`
	Active    bool           `json:"active" gorm:"default:true"`
	Role      Role           `json:"role" gorm:"default:'user'"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
