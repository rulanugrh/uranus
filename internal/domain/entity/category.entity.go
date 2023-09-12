package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string   `json:"name" form:"name" validate:"required"`
	Description string   `json:"desc" form:"desc" validate:"required"`
	Courses     []Course `json:"course" form:"course"`
}
