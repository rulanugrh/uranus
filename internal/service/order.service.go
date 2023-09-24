package service

import (
	"github.com/rulanugrh/uranus/internal/domain/entity"
	"github.com/rulanugrh/uranus/internal/domain/web"
	"github.com/rulanugrh/uranus/internal/repository/port"
	portServ "github.com/rulanugrh/uranus/internal/service/port"
	"github.com/rulanugrh/uranus/internal/utils"
)

type orderstruct struct {
	repository port.OrderInterfaceRepository
}

func NewOrderService(repo port.OrderInterfaceRepository) portServ.OrderInterfaceService {
	return &orderstruct{
		repository: repo,
	}
}

func(orderserv *orderstruct) CreateOrder(req entity.Order) (*web.ResponseOrder, error) {
	data, err := orderserv.repository.CreateOrder(req)
	if err != nil {
		return nil, err
	}

	response := utils.PrintResultOrder(*data)
	return &response, nil
}
func(ordeserv *orderstruct) FindByID(id uint) (*web.ResponseOrder, error) {
	data, err := ordeserv.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := utils.PrintResultOrder(*data)
	return &response, nil
}