package main

import (
	"github.com/rulanugrh/uranus/configs"
	"github.com/rulanugrh/uranus/internal/domain/entity"
	handler "github.com/rulanugrh/uranus/internal/http"
	"github.com/rulanugrh/uranus/internal/repository"
	"github.com/rulanugrh/uranus/internal/service"
	"github.com/rulanugrh/uranus/route"
)

func main() {
	db := configs.GetMysqlConn()
	db.AutoMigrate(&entity.Category{}, &entity.Course{}, &entity.Order{}, &entity.User{})

	userRepository := repository.NewUserRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)
	courseRepository := repository.NewCourseRepository(db)
	orderRepository := repository.NewOrderRepository(db)

	userService := service.NewUserService(userRepository)
	categoryService := service.NewCategoryServices(categoryRepository)
	courseService := service.NewCourseService(courseRepository)
	orderService := service.NewOrderService(orderRepository)

	userHandler := handler.NewUserHandler(userService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	courseHandler := handler.NewCourseHandler(courseService)
	orderHandler := handler.NewOrderHandler(orderService)

	route.Run(courseHandler, orderHandler, userHandler, categoryHandler)
}
