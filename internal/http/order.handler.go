package handler

import (
	"net/http"

	portHandler "github.com/rulanugrh/uranus/internal/http/port"
	"github.com/rulanugrh/uranus/internal/service/port"
)

type orderhandler struct {
	service port.OrderInterfaceService
}

func NewOrderHandler(serv port.OrderInterfaceService) portHandler.OrderInterfaceHTTP {
	return &orderhandler{
		service: serv,
	}
}

func(hnd *orderhandler) CreateOrder(w http.ResponseWriter, r *http.Request) {}

func(hnd *orderhandler) FindByID(w http.ResponseWriter, r *http.Request) {}