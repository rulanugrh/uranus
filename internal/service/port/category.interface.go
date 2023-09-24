package port

import (
	"github.com/rulanugrh/uranus/internal/domain/entity"
	"github.com/rulanugrh/uranus/internal/domain/web"
)

type CategoryInterfaceService interface {
	CreateCategory(req entity.Category) (*web.ResponseCreateCategory, error)
	FindByID(id uint) (*web.ResponseFindCategory, error)
	FindAll() ([]web.ResponseFindCategory, error)
}