package port

import "net/http"

type OrderInterfaceHTTP interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
	Checkout(w http.ResponseWriter, r *http.Request)
}
