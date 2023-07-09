package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Usucorreo      string `gorm:"not null:unique_index"`
	Usucontrasenia string `gorm:"not null"`
	Usutoken       string `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Category       []Category
}
