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

type categoryhandler struct {
	service port.CategoryInterfaceService
	validate *validator.Validate
}

func NewCategoryHandler(serv port.CategoryInterfaceService) portHandler.CategoryIntefaceHTTP {
	return &categoryhandler{
		service: serv,
		validate: validator.New(),
	}
}

func (hnd *categoryhandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var req entity.Category
	data, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(data, &req)
	errStruct := middleware.ValidateStruct(hnd.validate, req)
	if errStruct != nil {
		res := web.WebValidationError {
			Message: "cant create category",
			Error: errStruct,
		}
		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)

	}
	
	result, err := hnd.service.CreateCategory(req)
	if err != nil {
		res := web.ResponseFailure{
			Code:    http.StatusBadRequest,
			Message: "Cant create category",
		}

		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	res := web.ResponseSuccess{
		Code:    200,
		Message: "Success create category",
		Data:    result,
	}

	response, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (hnd *categoryhandler) FindByID(w http.ResponseWriter, r *http.Request) {
	getID := mux.Vars(r)
	paramsID := getID["id"]

	id, _ := strconv.Atoi(paramsID)

	result, err := hnd.service.FindByID(uint(id))
	if err != nil {
		res := web.ResponseFailure{
			Code:    http.StatusBadRequest,
			Message: "Cant find category by this id",
		}

		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	res := web.ResponseSuccess{
		Code:    http.StatusOK,
		Message: "Success find category",
		Data:    result,
	}

	response, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (hnd *categoryhandler) FindAll(w http.ResponseWriter, r *http.Request) {
	result, err := hnd.service.FindAll()
	if err != nil {
		res := web.ResponseFailure{
			Code:    http.StatusBadRequest,
			Message: "Cant find category",
		}

		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	res := web.ResponseSuccess{
		Code:    http.StatusOK,
		Message: "Success find category",
		Data:    result,
	}

	response, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
