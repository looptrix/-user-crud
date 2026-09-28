package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"user-crud/sqlx/model"
	"user-crud/sqlx/service"
)

// UserHandler обрабатывает HTTP-запросы пользователей.
type UserHandler struct {
	service *service.UserService
}

// NewUserHandler создаёт новый UserHandler.
func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// Create обрабатывает POST /users.
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "неверный JSON", http.StatusBadRequest)
		return
	}

	err = h.service.Create(&user)
	if err != nil {
		http.Error(w, "ошибка создания пользователя", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)
}

// GetAll обрабатывает GET /users.
func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "ошибка получения пользователей", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}

// GetByID обрабатывает GET /users/{id}.
func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r.URL.Path)
	if err != nil {
		http.Error(w, "неверный ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, "пользователь не найден", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

// Update обрабатывает PUT /users/{id}.
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r.URL.Path)
	if err != nil {
		http.Error(w, "неверный ID", http.StatusBadRequest)
		return
	}

	var user model.User

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "неверный JSON", http.StatusBadRequest)
		return
	}

	user.ID = id

	err = h.service.Update(&user)
	if err != nil {
		http.Error(w, "ошибка обновления пользователя", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

// Delete обрабатывает DELETE /users/{id}.
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r.URL.Path)
	if err != nil {
		http.Error(w, "неверный ID", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		http.Error(w, "ошибка удаления пользователя", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// getID получает ID из URL.
// Например: /users/5 → ID = 5.
func getID(path string) (uint, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")

	id, err := strconv.ParseUint(parts[len(parts)-1], 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}
