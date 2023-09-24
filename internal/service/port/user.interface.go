package port

import (
	"github.com/rulanugrh/uranus/internal/domain/entity"
	"github.com/rulanugrh/uranus/internal/domain/web"
)

type UserInterfaceService interface {
	CreateUser(req entity.User) (*web.ResponseUser, error)
	UpdateUser(id uint, req entity.User) (*web.ResponseUser, error)
	FindByID(id uint) (*web.ResponseUser, error)
	DeleteUser(id uint) error
}