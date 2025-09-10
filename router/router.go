package router

import (
	"CRUD_operation/Controller"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(userHandler *Controller.UserHandler) *mux.Router {
	r := mux.NewRouter()

	// Health check (optional)
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API is running"))
	}).Methods("GET")

	// CRUD routes
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	return r
}
