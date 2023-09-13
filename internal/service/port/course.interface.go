package port

import (
	"github.com/rulanugrh/uranus/internal/domain/entity"
	"github.com/rulanugrh/uranus/internal/domain/web"
)

type CourseInterfaceService interface {
	CreateCourse(req entity.Course) (*web.ResponseCreateCourse, error)
	FindCourse() ([]web.ResponseFindCourse, error)
	FindById(id uint) (*web.ResponseFindCourse, error)
	UpdateCourse(id uint, req entity.Course) (*web.ResponseFindCourse, error)
	DeleteCourse(id uint) error
}