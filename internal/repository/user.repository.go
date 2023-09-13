package repository

import (
	"context"
	"log"

	"github.com/rulanugrh/uranus/internal/domain/entity"
	"github.com/rulanugrh/uranus/internal/repository/port"
	"gorm.io/gorm"
)

type userstruct struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.UserInterfaceRepository {
	return &userstruct {
		DB: db,
	}
}

func (repo *userstruct) CreateUser(req entity.User) (*entity.User, error) {
	err := repo.DB.WithContext(context.Background()).Create(&req).Error
	if err != nil {
		log.Printf("Cant create user: %v", err)
		return nil, err
	}

	return &req, nil
}

func (repo *userstruct) FindByID(id uint) (*entity.User, error) {
	var req entity.User

	err := repo.DB.WithContext(context.Background()).Where("id = ?", id).Find(&req).Error
	if err != nil {
		log.Printf("cant find user by id: %v", err)
		return nil, err
	}

	return &req, nil
}

func (repo *userstruct) UpdateUser(id uint, req entity.User) (*entity.User, error) {
	var user entity.User

	err := repo.DB.WithContext(context.Background()).Model(&req).Where("id = ?").Updates(&user).Error
	if err != nil {
		log.Printf("cant update user by this id: %v", err)
		return nil, err
	}

	return &user, nil
}

func (repo *userstruct) DeleteUser(id uint) error {
	var user entity.User

	err := repo.DB.WithContext(context.Background()).Where("id = ?", id).Delete(&user).Error
	if err != nil {
		log.Printf("cant delete user by this id: %v", err)
		return err
	}

	return nil
}