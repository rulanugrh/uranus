package route

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rulanugrh/uranus/configs"
	"github.com/rulanugrh/uranus/internal/http/port"
)

func Run(course port.CourseInterfaceHTTP, order port.OrderInterfaceHTTP, user port.UserInterfaceHTTP, category port.CategoryIntefaceHTTP) {
	router := mux.NewRouter().StrictSlash(true)
	conf := configs.GetConfig()

	// endpoint course
	router.HandleFunc("/api/course", course.CreateCourse).Methods("POST")
	router.HandleFunc("/api/course/{id}", course.FindByID).Methods("GET")
	router.HandleFunc("/api/course/{id}", course.Update).Methods("PUT")
	router.HandleFunc("/api/course/{id}", course.Delete).Methods("DELETE")

	// endpoint order
	router.HandleFunc("/api/order", order.CreateOrder).Methods("POST")
	router.HandleFunc("/api/order/{id}", order.FindByID).Methods("GET")

	// endpoint user
	router.HandleFunc("/api/user", user.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", user.Update).Methods("PUT")
	router.HandleFunc("/api/user/{id}", user.FindByID).Methods("GET")
	router.HandleFunc("/api/user/{id}", user.Delete).Methods("DELETE")

	// endpoint category
	router.HandleFunc("/api/category", category.CreateCategory).Methods("POST")
	router.HandleFunc("/api/category", category.FindAll).Methods("GET")
	router.HandleFunc("/api/category/{id}", category.FindByID).Methods("GET")

	server := http.Server{
		Addr: conf.Server.Host + ":" + conf.Server.Port,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println("Cant connect to server")
	}
}
