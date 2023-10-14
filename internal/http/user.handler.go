package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

type userhandler struct {
	service  port.UserInterfaceService
	validate *validator.Validate
}

func NewUserHandler(serv port.UserInterfaceService) portHandler.UserInterfaceHTTP {
	return &userhandler{
		service:  serv,
		validate: validator.New(),
	}
}

func (hnd *userhandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req entity.User
	data, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(data, &req)
	password := middleware.HashPassword(req.Password)
	req.Password = password

	errStruct := middleware.ValidateStruct(hnd.validate, req)
	if errStruct != nil {
		res := web.WebValidationError{
			Message: "cant create user",
			Error:   errStruct,
		}
		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)

	}

	result, err := hnd.service.CreateUser(req)
	if err != nil {
		res := web.ResponseFailure{
			Code:    http.StatusBadRequest,
			Message: "Cant create user",
		}

		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	token, errToken := middleware.GenerateToken(req)
	if errToken != nil {
		log.Printf("Cant generate token because: %v", errToken)
	}

	res := web.ResponseSuccessAuth{
		Code:    200,
		Message: "Success create user",
		Data:    result,
		Token:   token,
	}

	response, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (hnd *userhandler) FindByID(w http.ResponseWriter, r *http.Request) {
	getID := mux.Vars(r)
	paramsID := getID["id"]

	id, _ := strconv.Atoi(paramsID)

	result, err := hnd.service.FindByID(uint(id))
	if err != nil {
		res := web.ResponseFailure{
			Code:    http.StatusBadRequest,
			Message: "Cant find user by this id",
		}

		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	res := web.ResponseSuccess{
		Code:    http.StatusOK,
		Message: "Success find user",
		Data:    result,
	}

	response, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (hnd *userhandler) Update(w http.ResponseWriter, r *http.Request) {
	var req entity.User
	data, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(data, &req)

	getID := mux.Vars(r)
	paramsID := getID["id"]

	id, _ := strconv.Atoi(paramsID)

	result, err := hnd.service.UpdateUser(uint(id), req)
	if err != nil {
		res := web.ResponseFailure{
			Code:    http.StatusBadRequest,
			Message: "Cant update user by this id",
		}

		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	res := web.ResponseSuccess{
		Code:    http.StatusOK,
		Message: "Success update user",
		Data:    result,
	}

	response, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (hnd *userhandler) Delete(w http.ResponseWriter, r *http.Request) {
	getID := mux.Vars(r)
	paramsID := getID["id"]

	id, _ := strconv.Atoi(paramsID)

	err := hnd.service.DeleteUser(uint(id))
	if err != nil {
		res := web.ResponseFailure{
			Code:    http.StatusBadRequest,
			Message: "Cant delete user by this id",
		}

		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	res := web.ResponseSuccess{
		Code:    http.StatusOK,
		Message: "Success delete user",
		Data:    "Data success delete",
	}

	response, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (hnd *userhandler) Login(w http.ResponseWriter, r *http.Request) {
	var req entity.User
	data, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(data, &req)

	result, err := hnd.service.FindByEmail(req.Email)
	if err != nil {
		res := web.ResponseFailure{
			Code:    http.StatusNotFound,
			Message: "cant login",
		}

		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	compare := []byte(req.Password)
	if errCheck := middleware.ComparePassword(req.Password, compare); errCheck != nil {
		res := web.ResponseFailure{
			Message: "Password not matched",
		}
		log.Printf("password not matched because: %v", err)
		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(response)
	}

	res := web.ResponseSuccess{
		Code:    http.StatusOK,
		Message: "success login",
		Data:    result,
	}
	response, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (hnd *userhandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req entity.User
	data, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(data, &req)

	result, err := hnd.service.FindByEmail(req.Email)
	if err != nil {
		res := web.ResponseFailure{
			Code:    http.StatusNotFound,
			Message: "cant refresh token",
		}

		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	token, errToken := middleware.GenerateToken(req)
	if errToken != nil {
		log.Printf("Cant generate token because: %v", errToken)
	}

	res := web.ResponseSuccessAuth{
		Code:    http.StatusOK,
		Message: "success refreh token ",
		Data:    result,
		Token:   token,
	}
	response, _ := json.Marshal(res)

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
