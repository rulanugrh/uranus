package port

import "github.com/rulanugrh/uranus/internal/domain/entity"

type UserInterfaceRepository interface {
	CreateUser(req entity.User) (*entity.User, error)
	FindByID(id uint) (*entity.User, error)
	UpdateUser(id uint, req entity.User) (*entity.User, error)
	DeleteUser(id uint) error
}