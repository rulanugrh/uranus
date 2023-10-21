package port

import "net/http"

type OrderInterfaceHTTP interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
	TestCheckout(w http.ResponseWriter, r *http.Request)
	History(w http.ResponseWriter, r *http.Request)
}
