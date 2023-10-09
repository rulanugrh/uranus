package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rulanugrh/uranus/internal/domain/entity"
	"github.com/rulanugrh/uranus/internal/domain/web"
	portHandler "github.com/rulanugrh/uranus/internal/http/port"
	"github.com/rulanugrh/uranus/internal/middleware"
	"github.com/rulanugrh/uranus/internal/service/port"
)

type orderhandler struct {
	service port.OrderInterfaceService
	validate *validator.Validate
}

func NewOrderHandler(serv port.OrderInterfaceService) portHandler.OrderInterfaceHTTP {
	return &orderhandler{
		service: serv,
		validate: validator.New(),
	}
}

func (hnd *orderhandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req entity.Order
	data, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(data, &req)
	errStruct := middleware.ValidateStruct(hnd.validate, req)
	if errStruct != nil {
		res := web.WebValidationError {
			Message: "cant create order",
			Error: errStruct,
		}
		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)

	}
	
	result, err := hnd.service.CreateOrder(req)
	if err != nil {
		res := web.ResponseFailure{
			Code:    http.StatusBadRequest,
			Message: "Cant create order",
		}

		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	res := web.ResponseSuccess{
		Code:    200,
		Message: "Success create order",
		Data:    result,
	}

	response, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (hnd *orderhandler) FindByID(w http.ResponseWriter, r *http.Request) {
	getID := mux.Vars(r)
	paramsID := getID["id"]

	id, _ := strconv.Atoi(paramsID)

	result, err := hnd.service.FindByID(uint(id))
	if err != nil {
		res := web.ResponseFailure{
			Code:    http.StatusBadRequest,
			Message: "Cant find order by this id",
		}

		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	res := web.ResponseSuccess{
		Code:    http.StatusOK,
		Message: "Success find order",
		Data:    result,
	}

	response, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
