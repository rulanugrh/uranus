package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Address  string `json:"address" form:"address" validate:"required"`
	Avatar   string `json:"avatar" form:"avatar"`
	Notelp   int    `json:"no_telp" form:"no_telp" validate:"required"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}
