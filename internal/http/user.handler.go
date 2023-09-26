package handler

import (
	"net/http"

	portHandler "github.com/rulanugrh/uranus/internal/http/port"
	"github.com/rulanugrh/uranus/internal/service/port"
)


type userhandler struct {
	service port.UserInterfaceService
}

func NewUserHandler(serv port.UserInterfaceService) portHandler.UserInterfaceHTTP {
	return &userhandler{
		service: serv,
	}
}

func(hnd *userhandler) CreateUser(w http.ResponseWriter, r *http.Request) {}

func(hnd *userhandler) FindByID(w http.ResponseWriter, r *http.Request) {}

func(hnd *userhandler) Update(w http.ResponseWriter, r *http.Request) {}

func(hnd *userhandler) Delete(w http.ResponseWriter, r *http.Request) {}