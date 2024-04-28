package createUser

import (
	"encoding/json"
	"net/http"
)

type (
	userCreateService interface {
		Create(name, surname string) error
	}

	UserCreateController struct {
		service userCreateService
	}

	user struct {
		Name    string `json:"name"`
		Surname string `json:"surname"`
	}
)

func NewController(service userCreateService) *UserCreateController {
	return &UserCreateController{
		service: service,
	}
}

func (c *UserCreateController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser user

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = c.service.Create(newUser.Name, newUser.Surname)
	if err != nil {
		http.Error(w, "Error on creating user / controller layer", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": "User is created successfully in phone book"}
	json.NewEncoder(w).Encode(response)
}
