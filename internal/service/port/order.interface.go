package port

import (
	"github.com/rulanugrh/uranus/internal/domain/entity"
	"github.com/rulanugrh/uranus/internal/domain/web"
)

type OrderInterfaceService interface {
	CreateOrder(req entity.Order) (*web.ResponseOrder, error)
	FindByID(id uint) (*web.ResponseOrder, error)
}