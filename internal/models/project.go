package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model

	Name string `gorm:"not null"`

	GitURL string `gorm:"not null"`

	Branch string `gorm:"default:main"`

	Description string

	Status string `gorm:"default:created"`

	UserID uint
}
