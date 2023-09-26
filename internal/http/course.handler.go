package handler

import (
	"net/http"

	portHandler "github.com/rulanugrh/uranus/internal/http/port"
	"github.com/rulanugrh/uranus/internal/service/port"
)


type coursehandler struct {
	service port.CourseInterfaceService
}

func NewCourseHandler(serv port.CourseInterfaceService) portHandler.CourseInterfaceHTTP {
	return &coursehandler{
		service: serv,
	}
}

func(hnd *coursehandler) CreateCourse(w http.ResponseWriter, r *http.Request) {}

func(hnd *coursehandler) FindByID(w http.ResponseWriter, r *http.Request) {}

func(hnd *coursehandler) Update(w http.ResponseWriter, r *http.Request) {}

func(hnd *coursehandler) FindAll(w http.ResponseWriter, r *http.Request) {}

func(hnd *coursehandler) Delete(w http.ResponseWriter, r *http.Request) {}