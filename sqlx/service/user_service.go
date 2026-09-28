package service

import (
	"user-crud/sqlx/model"
	"user-crud/sqlx/repository"
)

// UserService содержит логику работы с пользователями.
type UserService struct {
	repo *repository.UserRepository
}

// NewUserService создаёт новый UserService.
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// Create создаёт пользователя.
func (s *UserService) Create(user *model.User) error {
	return s.repo.Create(user)
}

// GetByID получает пользователя по ID.
func (s *UserService) GetByID(id uint) (*model.User, error) {
	return s.repo.GetByID(id)
}

// GetAll получает всех пользователей.
func (s *UserService) GetAll() ([]model.User, error) {
	return s.repo.GetAll()
}

// Update изменяет пользователя.
func (s *UserService) Update(user *model.User) error {
	return s.repo.Update(user)
}

// Delete удаляет пользователя.
func (s *UserService) Delete(id uint) error {
	return s.repo.Delete(id)
}
