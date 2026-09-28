package service

import (
	"context"

	"user-crud/sqlc/generated"
	"user-crud/sqlc/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(
	ctx context.Context,
	params generated.CreateUserParams,
) (generated.User, error) {
	return s.repo.Create(ctx, params)
}

func (s *UserService) GetByID(
	ctx context.Context,
	id int64,
) (generated.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) GetAll(
	ctx context.Context,
) ([]generated.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *UserService) Update(
	ctx context.Context,
	params generated.UpdateUserParams,
) (generated.User, error) {
	return s.repo.Update(ctx, params)
}

func (s *UserService) Delete(
	ctx context.Context,
	id int64,
) error {
	return s.repo.Delete(ctx, id)
}
