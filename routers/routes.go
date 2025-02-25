package routers

import (
	"Codimite_Assignment/internal/handlers"
	"Codimite_Assignment/pkg/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func SetRoutes() *mux.Router {
	r := mux.NewRouter()


	r.Handle("/user",middleware.AuthMiddleware(http.HandlerFunc(handlers.UpdateUser))).Methods(http.MethodPut)
	r.Handle("/user",middleware.AuthMiddleware(http.HandlerFunc(handlers.DeleteUser))).Methods(http.MethodDelete)
	r.HandleFunc("/users",handlers.GetAllUsers).Methods(http.MethodGet)

	r.HandleFunc("/register",handlers.RegisterUser).Methods(http.MethodPost)
	r.HandleFunc("/login",handlers.LoginUser).Methods(http.MethodPost)

	r.HandleFunc("/products",handlers.CreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/products/{id}",handlers.UpdateProduct).Methods(http.MethodPut)
	r.HandleFunc("/products",handlers.GetAllProducts).Methods(http.MethodGet)
	r.HandleFunc("/products/{id}",handlers.DeleteProduct).Methods(http.MethodDelete)

	r.HandleFunc("/orders",handlers.CreateOrder).Methods(http.MethodPost)
	r.HandleFunc("/orders/{id}",handlers.UpdateOrder).Methods(http.MethodPut)
	r.HandleFunc("/orders",handlers.GetAllOrder).Methods(http.MethodGet)
	r.HandleFunc("/orders/{id}",handlers.DeleteOrder).Methods(http.MethodDelete)


	return r
}