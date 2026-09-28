package repository

import (
	"user-crud/sqlx/model"

	"github.com/jmoiron/sqlx"
)

// UserRepository работает с таблицей users
// напрямую через SQL-запросы.
type UserRepository struct {
	db *sqlx.DB
}

// NewUserRepository создаёт repository.
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create создаёт нового пользователя.
func (r *UserRepository) Create(user *model.User) error {
	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	return r.db.QueryRowx(
		query,
		user.Name,
		user.Email,
		user.Password,
	).Scan(&user.ID)
}

// GetByID получает одного пользователя по ID.
func (r *UserRepository) GetByID(id uint) (*model.User, error) {
	var user model.User

	query := `
		SELECT id, name, email, password
		FROM users
		WHERE id = $1
	`

	err := r.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetAll получает всех пользователей.
func (r *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User

	query := `
		SELECT id, name, email, password
		FROM users
		ORDER BY id
	`

	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Update изменяет пользователя.
func (r *UserRepository) Update(user *model.User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2, password = $3
		WHERE id = $4
	`

	_, err := r.db.Exec(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.ID,
	)

	return err
}

// Delete удаляет пользователя.
func (r *UserRepository) Delete(id uint) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	_, err := r.db.Exec(query, id)

	return err
}
