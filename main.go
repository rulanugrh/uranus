package main

import (
	"github.com/rulanugrh/uranus/configs"
	"github.com/rulanugrh/uranus/internal/domain/entity"
	handler "github.com/rulanugrh/uranus/internal/http"
	"github.com/rulanugrh/uranus/internal/repository"
	"github.com/rulanugrh/uranus/internal/service"
	"github.com/rulanugrh/uranus/route"
	payment "github.com/rulanugrh/uranus/third_party/midtrans"
)

func main() {
	configs.SetupMidtransSandbox()
	db := configs.GetMysqlConn()
	db.AutoMigrate(&entity.Category{}, &entity.Course{}, &entity.Order{}, &entity.User{})

	userRepository := repository.NewUserRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)
	courseRepository := repository.NewCourseRepository(db)
	orderRepository := repository.NewOrderRepository(db)
	paymentRepository := repository.NewPaymentRepository(db)

	userService := service.NewUserService(userRepository)
	categoryService := service.NewCategoryServices(categoryRepository)
	courseService := service.NewCourseService(courseRepository)
	orderService := service.NewOrderService(orderRepository)

	userHandler := handler.NewUserHandler(userService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	courseHandler := handler.NewCourseHandler(courseService)

	paymentMethod := payment.NewPayment(paymentRepository ,userRepository, orderRepository, courseRepository)

	orderHandler := handler.NewOrderHandler(orderService, paymentMethod)

	route.Run(courseHandler, orderHandler, userHandler, categoryHandler)
}
