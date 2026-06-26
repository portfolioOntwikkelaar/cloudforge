package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model

	Name string

	GitURL string

	UserID uint
}
