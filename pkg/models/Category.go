package models

import "time"

type Category struct {
	Catname   string `gorm:"not null"`
	Catcolor  string `gorm:"null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint
}
