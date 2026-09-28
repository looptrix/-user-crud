package model

// User — модель пользователя.
// GORM использует эту структуру,
// чтобы создать и изменять таблицу users в PostgreSQL.
type User struct {
	ID       uint   `gorm:"primaryKey"` // уникальный ID пользователя
	Name     string // имя пользователя
	Email    string // email пользователя
	Password string // пароль пользователя
}
