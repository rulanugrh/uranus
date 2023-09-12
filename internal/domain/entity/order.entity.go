package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID    int    `json:"user_id" form:"user_id" validate:"required"`
	CourseID  int    `json:"course_id" form:"course_id" validate:"required"`
	UserMod   User   `json:"user" form:"user" gorm:"foreignKey:UserID;reference:ID"`
	CourseMod Course `json:"course" form:"course" gorm:"foreignKey:CourseID;reference:ID"`
}
