package handlers


import (
	"encoding/json"
	"net/http"
	"simple-crud/internal/models"
	"simple-crud/internal/repository"
	"time"
	"fmt"
	"database/sql"
    "github.com/gorilla/mux"
    "strconv"
)

type UserHandler struct {
	userRepository *repository.UserRepository
}

func NewUserHandler(userRepository *repository.UserRepository) *UserHandler {
	return &UserHandler{userRepository: userRepository}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Email == "" {
		http.Error(w, "Username and email are required", http.StatusBadRequest)
		return
	}


	fmt.Println(user)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := h.userRepository.CreateNewUser(&user); err!= nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.ParseInt(vars["id"], 10, 64)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    user, err := h.userRepository.GetUserById(id)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "User not found", http.StatusNotFound)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}


func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.ParseInt(vars["id"], 10, 64)
    if err!= nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err!= nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    user.ID = id
    user.UpdatedAt = time.Now()
    if err := h.userRepository.UpdateUser(&user); err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.ParseInt(vars["id"], 10, 64)
    if err!= nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }
    if err := h.userRepository.DeleteUser(id); err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
    users, err := h.userRepository.GetAllUsers()
    if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}