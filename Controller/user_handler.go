package Controller

import (
	"encoding/json"
	"net/http"

	"CRUD_operation/models"
	"CRUD_operation/service"
	"CRUD_operation/validator"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.Schema
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validator.ValidateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := h.userService.CreateUser(user.Name, user.Age, user.Email, user.College, user.Branch, user.Phone, user.Address, user.Language, user.Projects, user.About, user.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	user, err := h.userService.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var user models.Schema
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedUser, err := h.userService.UpdateUser(id, user.Name, user.Age, user.Email, user.College, user.Branch, user.Phone, user.Address, user.Language, user.Projects, user.About, user.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User updated successfully",
		"user":    updatedUser,
	})
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	err := h.userService.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User deleted successfully",
	})

}
