package route

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rulanugrh/uranus/configs"
	"github.com/rulanugrh/uranus/internal/http/port"
	"github.com/rulanugrh/uranus/internal/middleware"
)

func Run(course port.CourseInterfaceHTTP, order port.OrderInterfaceHTTP, user port.UserInterfaceHTTP, category port.CategoryIntefaceHTTP) {
	router := mux.NewRouter().StrictSlash(true)
	conf := configs.GetConfig()

	router.HandleFunc("/user/auth", user.CreateUser).Methods("POST")
	router.HandleFunc("/user/login", user.Login).Methods("POST")
	router.Use(middleware.CommonMiddleware)

	routerHandler := router.PathPrefix("/api").Subrouter()
	routerHandler.Use(middleware.JWTVerify)
	routerHandler.Use(middleware.CommonMiddleware)

	// endpoint course
	routerHandler.HandleFunc("/course", course.CreateCourse).Methods("POST")
	routerHandler.HandleFunc("/course/{id}", course.FindByID).Methods("GET")
	routerHandler.HandleFunc("/course/{id}", course.Update).Methods("PUT")
	routerHandler.HandleFunc("/course/{id}", course.Delete).Methods("DELETE")

	// endpoint order
	routerHandler.HandleFunc("/order", order.CreateOrder).Methods("POST")
	routerHandler.HandleFunc("/order/checkout/{id}", order.TestCheckout).Methods("POST")
	routerHandler.HandleFunc("/order/{id}", order.FindByID).Methods("GET")
	routerHandler.HandleFunc("/order/history/{id}", order.History).Methods("GET")

	// endpoint user

	routerHandler.HandleFunc("/user/{id}", user.Update).Methods("PUT")
	routerHandler.HandleFunc("/user/{id}", user.FindByID).Methods("GET")
	routerHandler.HandleFunc("/user/{id}", user.Delete).Methods("DELETE")

	// endpoint category
	routerHandler.HandleFunc("/category", category.CreateCategory).Methods("POST")
	routerHandler.HandleFunc("/category", category.FindAll).Methods("GET")
	routerHandler.HandleFunc("/category/{id}", category.FindByID).Methods("GET")

	server := http.Server{
		Addr: conf.Server.Host + ":" + conf.Server.Port,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println("Cant connect to server")
	}
}
