package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rulanugrh/uranus/internal/domain/entity"
	"github.com/rulanugrh/uranus/internal/domain/web"
	portHandler "github.com/rulanugrh/uranus/internal/http/port"
	"github.com/rulanugrh/uranus/internal/middleware"
	"github.com/rulanugrh/uranus/internal/service/port"
	"github.com/rulanugrh/uranus/third_party/monitoring"
)

type coursehandler struct {
	service  port.CourseInterfaceService
	validate *validator.Validate
}

func NewCourseHandler(serv port.CourseInterfaceService) portHandler.CourseInterfaceHTTP {
	return &coursehandler{
		service:  serv,
		validate: validator.New(),
	}
}

func (hnd *coursehandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var req entity.Course
	data, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(data, &req)
	errStruct := middleware.ValidateStruct(hnd.validate, req)
	if errStruct != nil {
		res := web.WebValidationError{
			Message: "cant create course",
			Error:   errStruct,
		}
		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)

	}

	tracing := monitoring.StartTracing(r.Context(), "Handle Create Course")
	time.Sleep(time.Second / 2)
	tracing.Finish()

	result, err := hnd.service.CreateCourse(req)
	if err != nil {
		res := web.ResponseFailure{
			Code:    http.StatusBadRequest,
			Message: "Cant create course",
		}

		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	res := web.ResponseSuccess{
		Code:    200,
		Message: "Success create course",
		Data:    result,
	}

	response, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (hnd *coursehandler) FindByID(w http.ResponseWriter, r *http.Request) {
	getID := mux.Vars(r)
	paramsID := getID["id"]

	id, _ := strconv.Atoi(paramsID)

	tracing := monitoring.StartTracing(r.Context(), "Handle Find Course")
	time.Sleep(time.Second / 2)
	tracing.Finish()

	result, err := hnd.service.FindById(uint(id))
	if err != nil {
		res := web.ResponseFailure{
			Code:    http.StatusBadRequest,
			Message: "Cant find course by this id",
		}

		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	res := web.ResponseSuccess{
		Code:    http.StatusOK,
		Message: "Success find course",
		Data:    result,
	}

	response, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (hnd *coursehandler) Update(w http.ResponseWriter, r *http.Request) {
	var req entity.Course
	data, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(data, &req)

	getID := mux.Vars(r)
	paramsID := getID["id"]

	id, _ := strconv.Atoi(paramsID)

	tracing := monitoring.StartTracing(r.Context(), "Handle Update Course")
	time.Sleep(time.Second / 2)
	tracing.Finish()

	result, err := hnd.service.UpdateCourse(uint(id), req)
	if err != nil {
		res := web.ResponseFailure{
			Code:    http.StatusBadRequest,
			Message: "Cant update course by this id",
		}

		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	res := web.ResponseSuccess{
		Code:    http.StatusOK,
		Message: "Success update course",
		Data:    result,
	}

	response, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (hnd *coursehandler) FindAll(w http.ResponseWriter, r *http.Request) {
	tracing := monitoring.StartTracing(r.Context(), "Handle FindAll Course")
	time.Sleep(time.Second / 2)
	tracing.Finish()

	result, err := hnd.service.FindCourse()
	if err != nil {
		res := web.ResponseFailure{
			Code:    http.StatusBadRequest,
			Message: "Cant find course",
		}

		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	res := web.ResponseSuccess{
		Code:    http.StatusOK,
		Message: "Success find course",
		Data:    result,
	}

	response, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (hnd *coursehandler) Delete(w http.ResponseWriter, r *http.Request) {
	getID := mux.Vars(r)
	paramsID := getID["id"]

	id, _ := strconv.Atoi(paramsID)

	tracing := monitoring.StartTracing(r.Context(), "Handle Delete Course")
	time.Sleep(time.Second / 2)
	tracing.Finish()

	err := hnd.service.DeleteCourse(uint(id))
	if err != nil {
		res := web.ResponseFailure{
			Code:    http.StatusBadRequest,
			Message: "Cant delete course by this id",
		}

		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	res := web.ResponseSuccess{
		Code:    http.StatusOK,
		Message: "Success delete course",
		Data:    "Data success delete",
	}

	response, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
