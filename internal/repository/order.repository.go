package repository

import (
	"context"
	"log"

	"github.com/rulanugrh/uranus/internal/domain/entity"
	"github.com/rulanugrh/uranus/internal/repository/port"
	"gorm.io/gorm"
)

type orderstruct struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) port.OrderInterfaceRepository {
	return &orderstruct{
		DB: db,
	}
}

func (repo *orderstruct) CreateOrder(req entity.Order) (*entity.Order, error) {
	err := repo.DB.WithContext(context.Background()).Create(&req).Error
	if err != nil {
		log.Printf("cant crete order, because %v", err)
		return nil, err
	}

	return &req, nil
}

func (repo *orderstruct) FindByID(id uint) (*entity.Order, error) {
	var req entity.Order

	err := repo.DB.WithContext(context.Background()).Preload("UserMod").Preload("CourseMod").Where("id = ?", id).Find(&req).Error

	if err != nil {
		log.Printf("cant find order: %v", err)
		return nil, err
	}

	return &req, nil
}