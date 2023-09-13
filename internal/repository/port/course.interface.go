package port

import "github.com/rulanugrh/uranus/internal/domain/entity"

type CourseInterfaceRepository interface {
	CreateCourse(req entity.Course) (*entity.Course, error)
	FindCourse() ([]entity.Course, error)
	FindById(id uint) (*entity.Course, error)
	UpdateCourse(id uint, req entity.Course) (*entity.Course, error)
	DeleteCourse(id uint) error
}
