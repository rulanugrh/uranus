package handler

import (
	"net/http"

	portHandler "github.com/rulanugrh/uranus/internal/http/port"
	"github.com/rulanugrh/uranus/internal/service/port"
)

type categoryhandler struct {
	service port.CategoryInterfaceService
}

func NewCategoryHandler(serv port.CategoryInterfaceService) portHandler.CategoryIntefaceHTTP {
	return &categoryhandler{
		service: serv,
	}
}

func(hnd *categoryhandler) CreateCategory(w http.ResponseWriter, r *http.Request) {}

func(hnd *categoryhandler) FindByID(w http.ResponseWriter, r *http.Request) {}

func(hnd *categoryhandler) FindAll(w http.ResponseWriter, r *http.Request) {}