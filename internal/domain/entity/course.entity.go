package entity

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Name           string    `json:"name" form:"name" validate:"required"`
	CategoryID     uint      `json:"category_id" form:"category_id" validate:"required"`
	Categories     Category  `json:"category" form:"category" gorm:"foreignKey:CategoryID;reference:ID"`
	Participant    []Order   `json:"participant" form:"participant"`
	Price          int       `json:"price" form:"price" validate:"required"`
	MaxParticipant int       `json:"max_participant" form:"max_participant" validate:"required"`
	Waktu          time.Time `json:"waktu" form:"waktu" validate:"required"`
}
