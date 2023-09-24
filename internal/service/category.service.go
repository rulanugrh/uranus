package service

import (
	"github.com/rulanugrh/uranus/internal/domain/entity"
	"github.com/rulanugrh/uranus/internal/domain/web"
	"github.com/rulanugrh/uranus/internal/repository/port"
	portServ "github.com/rulanugrh/uranus/internal/service/port"
	"github.com/rulanugrh/uranus/internal/utils"
)
type categorystruct struct {
	repository port.CategoryInterfaceRepository
}

func NewCategoryServices(repo port.CategoryInterfaceRepository) portServ.CategoryInterfaceService {
	return &categorystruct{
		repository: repo,
	}
}

func(cateserv *categorystruct) CreateCategory(req entity.Category) (*web.ResponseCreateCategory, error) {
	data, err := cateserv.repository.CreateCategory(req)
	if err != nil {
		return nil, err
	}

	response := utils.PrintResultCreateCategory(*data)
	return &response, nil
}

func(cateserv *categorystruct) FindByID(id uint) (*web.ResponseFindCategory, error) {
	data, err := cateserv.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := utils.PrintResultCategoryByID(*data)
	return &response, nil
}

func(cateserv *categorystruct) FindAll() ([]web.ResponseFindCategory, error) {
	data, err := cateserv.repository.FindAll()
	if err != nil {
		return []web.ResponseFindCategory{}, nil
	}

	response := utils.PrintResultCategory(data)
	return response, nil
}