package repository

import (
	"context"
	"log"

	"github.com/rulanugrh/uranus/internal/domain/entity"
	"github.com/rulanugrh/uranus/internal/repository/port"
	"gorm.io/gorm"
)

type coursestruct struct {
	DB *gorm.DB
}

func NewCourseRepository(db *gorm.DB) port.CourseInterfaceRepository {
	return &coursestruct{DB: db}
}

func (repo *coursestruct) CreateCourse(req entity.Course) (*entity.Course, error ) {
	err := repo.DB.WithContext(context.Background()).Create(&req).Error
	if err != nil {
		log.Printf("Cant create course to db because: %v", err)
		return nil, err
	}

	return &req, nil
}

func (repo *coursestruct) FindCourse() ([]entity.Course, error) {
	var courses []entity.Course
	
	err := repo.DB.WithContext(context.Background()).Preload("Category").Preload("Participant.User").Find(&courses).Error
	if err != nil {
		log.Printf("Cant find course because: %v", err)
		return nil, err
	}

	return courses, nil
}

func (repo *coursestruct) FindById(id uint) (*entity.Course, error) {
	var course entity.Course

	err := repo.DB.WithContext(context.Background()).Preload("Category").Preload("Participant.User").Where("id = ?", id).Find(&course).Error
	if err != nil {
		log.Printf("Cant find course because: %v", err)
		return nil, err
	}

	return &course, nil
}

func (repo *coursestruct) UpdateCourse(id uint, req entity.Course) (*entity.Course, error) {
	var course entity.Course
	err := repo.DB.WithContext(context.Background()).Model(&req).Where("id = ?", id).Updates(&course).Error
	if err != nil {
		log.Printf("Cant update course, because: %v", err)
		return nil, err
	}

	return &course, nil
}

func (repo *coursestruct) DeleteCourse(id uint) error {
	var course entity.Course

	err := repo.DB.WithContext(context.Background()).Where("id = ?", id).Delete(&course).Error
	if err != nil {
		log.Printf("Cant delete course, because: %v", err)
		return err
	}
	
	return nil
}