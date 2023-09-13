package port

import "github.com/rulanugrh/uranus/internal/domain/entity"

type CategoryInterfaceRepository interface {
	CreateCategory(req entity.Category) (*entity.Category, error)
	FindByID(id uint) (*entity.Category, error)
	FindAll() ([]entity.Category, error)
}