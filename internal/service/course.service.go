package service

import (
	"github.com/rulanugrh/uranus/internal/domain/entity"
	"github.com/rulanugrh/uranus/internal/domain/web"
	"github.com/rulanugrh/uranus/internal/repository/port"
	portServ "github.com/rulanugrh/uranus/internal/service/port"
	"github.com/rulanugrh/uranus/internal/utils"
)
type coursestruct struct {
	repository port.CourseInterfaceRepository
}

func NewCourseService(repo port.CourseInterfaceRepository) portServ.CourseInterfaceService {
	return &coursestruct{
		repository: repo,
	}
}

func (serv *coursestruct) CreateCourse(req entity.Course) (*web.ResponseCreateCourse, error) {
	data, err := serv.repository.CreateCourse(req)
	if err != nil {
		return nil, err
	}

	result := utils.PrintResultCreateCourse(*data)

	return &result, nil
}

func(serv *coursestruct) FindCourse() ([]web.ResponseFindCourse, error) {
	data, err := serv.repository.FindCourse()
	if err != nil {
		return []web.ResponseFindCourse{}, err
	}


	result := utils.PrintResultAllCourse(data)
	return result, nil
}

func(serv *coursestruct) FindById(id uint) (*web.ResponseFindCourse, error) {
	data, err := serv.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	result := utils.PrintResultCourse(*data)
	return &result, nil
}

func(serv *coursestruct) UpdateCourse(id uint, req entity.Course) (*web.ResponseFindCourse, error) {
	data, err := serv.repository.UpdateCourse(id, req)
	if err != nil {
		return nil, err
	}

	result := utils.PrintResultCourse(*data)
	return &result, nil
}

func(serv *coursestruct) DeleteCourse(id uint) error {
	err := serv.repository.DeleteCourse(id)
	if err != nil {
		return err
	}

	return nil
}