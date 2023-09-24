package service

import (
	"github.com/rulanugrh/uranus/internal/domain/entity"
	"github.com/rulanugrh/uranus/internal/domain/web"
	"github.com/rulanugrh/uranus/internal/repository/port"
	portServ "github.com/rulanugrh/uranus/internal/service/port"
	"github.com/rulanugrh/uranus/internal/utils"
)
type userstruct struct {
	repository port.UserInterfaceRepository
}

func NewUserService(repo port.UserInterfaceRepository) portServ.UserInterfaceService {
	return &userstruct{
		repository: repo,
	}
}
func(userserv *userstruct) CreateUser(req entity.User) (*web.ResponseUser, error) {
	data, err := userserv.repository.CreateUser(req)
	if err != nil {
		return nil, err
	}

	response := utils.PrintResultUser(*data)
	return &response, nil
}

func(userserv *userstruct) UpdateUser(id uint, req entity.User) (*web.ResponseUser, error) {
	data, err := userserv.repository.UpdateUser(id, req)
	if err != nil {
		return nil, err
	}

	response := utils.PrintResultUser(*data)
	return &response, nil
}

func(userserv *userstruct) FindByID(id uint) (*web.ResponseUser, error) {
	data, err := userserv.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := utils.PrintResultUser(*data)
	return &response, nil
}

func(userserv *userstruct) DeleteUser(id uint) error {
	err := userserv.repository.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}