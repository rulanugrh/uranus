package port

import "net/http"

type CategoryIntefaceHTTP interface {
	CreateCategory(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
}