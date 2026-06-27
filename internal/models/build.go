package models

import (
	"gorm.io/gorm"
)

type Build struct {
	gorm.Model
	ProjectID uint
	Status    string
	Log       string
}
