package service

import (
	"user-crud/internal/cache"
	"user-crud/internal/model"
	"user-crud/internal/repository"
)

// CachedUserService — сервис пользователей с cache.
type CachedUserService struct {
	repo  *repository.UserRepository
	cache *cache.Cache
}

// NewCachedUserService создаёт сервис с cache.
func NewCachedUserService(
	repo *repository.UserRepository,
	cache *cache.Cache,
) *CachedUserService {
	return &CachedUserService{
		repo:  repo,
		cache: cache,
	}
}

// Create создаёт пользователя и сохраняет его в cache.
func (s *CachedUserService) Create(user *model.User) error {
	err := s.repo.Create(user)
	if err != nil {
		return err
	}

	s.cache.Set(user.ID, user)

	return nil
}

// GetByID сначала проверяет cache.
// Если пользователя нет — получает его из PostgreSQL.
func (s *CachedUserService) GetByID(id uint) (*model.User, error) {
	if value, ok := s.cache.Get(id); ok {
		return value.(*model.User), nil
	}

	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	s.cache.Set(id, user)

	return user, nil
}

// GetAll получает всех пользователей из PostgreSQL.
func (s *CachedUserService) GetAll() ([]model.User, error) {
	return s.repo.GetAll()
}

// Update обновляет пользователя в PostgreSQL и cache.
func (s *CachedUserService) Update(user *model.User) error {
	err := s.repo.Update(user)
	if err != nil {
		return err
	}

	s.cache.Set(user.ID, user)

	return nil
}

// Delete удаляет пользователя из PostgreSQL и cache.
func (s *CachedUserService) Delete(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	s.cache.Delete(id)

	return nil
}
