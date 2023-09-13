package repository

import (
	"context"
	"log"

	"github.com/rulanugrh/uranus/internal/domain/entity"
	"github.com/rulanugrh/uranus/internal/repository/port"
	"gorm.io/gorm"
)

type categorystruct struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) port.CategoryInterfaceRepository {
	return &categorystruct{
		DB: db,
	}
}

func (repo *categorystruct) CreateCategory(req entity.Category) (*entity.Category, error) {
	err := repo.DB.WithContext(context.Background()).Create(&req).Error

	if err != nil {
		log.Printf("Cant create category: %v", err)
		return nil, err
	}

	return &req, nil
}

func (repo *categorystruct) FindAll() ([]entity.Category, error) {
	var categories []entity.Category

	err := repo.DB.WithContext(context.Background()).Preload("Courses").Find(&categories).Error
	if err != nil {
		log.Printf("Cant find category: %v", err)
		return []entity.Category{}, err
	}

	return categories, nil
}

func (repo *categorystruct) FindByID(id uint) (*entity.Category, error) {
	var category entity.Category

	err := repo.DB.WithContext(context.Background()).Preload("Courses").Where("id = ?", id).Find(&category).Error
	if err != nil {
		log.Printf("cant find category: %v", err)
		return nil, err
	}

	return &category, nil
}
