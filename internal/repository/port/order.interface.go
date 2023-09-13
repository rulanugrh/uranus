package port

import "github.com/rulanugrh/uranus/internal/domain/entity"

type OrderInterfaceRepository interface {
	CreateOrder(req entity.Order) (*entity.Order, error)
	FindByID(id uint) (*entity.Order, error)
}