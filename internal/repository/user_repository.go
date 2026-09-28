package repository

import (
	"user-crud/internal/model"

	"gorm.io/gorm"
)

// UserRepository — объект, через который работаем
// с таблицей users в базе данных.
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository создаёт новый UserRepository.
// Получаем готовое подключение к базе db
// и сохраняем его внутри UserRepository.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create создаёт нового пользователя в базе данных.
func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// GetByID получает пользователя по его ID.
func (r *UserRepository) GetByID(id uint) (*model.User, error) {
	var user model.User

	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetAll получает всех пользователей.
func (r *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User

	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Update изменяет данные пользователя.
func (r *UserRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

// Delete удаляет пользователя по ID.
func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}
