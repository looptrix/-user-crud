package repository

import (
	"context"

	"user-crud/sqlc/generated"
)

type UserRepository struct {
	queries *generated.Queries
}

func NewUserRepository(queries *generated.Queries) *UserRepository {
	return &UserRepository{queries: queries}
}

func (r *UserRepository) Create(
	ctx context.Context,
	params generated.CreateUserParams,
) (generated.User, error) {
	return r.queries.CreateUser(ctx, params)
}

func (r *UserRepository) GetByID(
	ctx context.Context,
	id int64,
) (generated.User, error) {
	return r.queries.GetUser(ctx, id)
}

func (r *UserRepository) GetAll(
	ctx context.Context,
) ([]generated.User, error) {
	return r.queries.GetUsers(ctx)
}

func (r *UserRepository) Update(
	ctx context.Context,
	params generated.UpdateUserParams,
) (generated.User, error) {
	return r.queries.UpdateUser(ctx, params)
}

func (r *UserRepository) Delete(
	ctx context.Context,
	id int64,
) error {
	return r.queries.DeleteUser(ctx, id)
}
