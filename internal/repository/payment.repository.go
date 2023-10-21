package repository

import (
	"context"

	"github.com/rulanugrh/uranus/internal/domain/entity"
	"github.com/rulanugrh/uranus/internal/repository/port"
	"gorm.io/gorm"
)

type paymentstruct struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) port.PaymentRepository {
	return &paymentstruct{
		DB: DB,
	}
}

func (repo *paymentstruct) Save(req entity.PaymentSandbox) (*entity.PaymentSandbox, error) {
	err := repo.DB.WithContext(context.Background()).Create(req).Error
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func (repo *paymentstruct) History(id uint) (*entity.PaymentSandbox, error) {
	var req entity.PaymentSandbox
	err := repo.DB.WithContext(context.Background()).Where("id = ?", id).Find(&req).Error
	if err != nil {
		return nil, err
	}

	return &req, nil
}
