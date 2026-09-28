package main

import (
	"fmt"
	"log"
	"net/http"

	"user-crud/sqlx/handler"
	"user-crud/sqlx/repository"
	"user-crud/sqlx/service"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// Данные для подключения к PostgreSQL.
	dsn := "host=localhost user=postgres password=postgres dbname=user_crud port=5432 sslmode=disable"

	// Подключаемся к PostgreSQL через sqlx.
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal("ошибка подключения к PostgreSQL:", err)
	}

	// Проверяем, что соединение с базой работает.
	err = db.Ping()
	if err != nil {
		log.Fatal("PostgreSQL не отвечает:", err)
	}

	fmt.Println("PostgreSQL подключён через sqlx!")

	// Создаём repository.
	userRepository := repository.NewUserRepository(db)

	// Создаём service.
	userService := service.NewUserService(userRepository)

	// Создаём handler.
	userHandler := handler.NewUserHandler(userService)

	// POST /users — создать пользователя.
	// GET /users — получить всех пользователей.
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userHandler.Create(w, r)

		case http.MethodGet:
			userHandler.GetAll(w, r)

		default:
			http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})

	// GET /users/{id} — получить пользователя.
	// PUT /users/{id} — изменить пользователя.
	// DELETE /users/{id} — удалить пользователя.
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetByID(w, r)

		case http.MethodPut:
			userHandler.Update(w, r)

		case http.MethodDelete:
			userHandler.Delete(w, r)

		default:
			http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("SQLX сервер запущен на http://localhost:8081")

	// Используем порт 8081,
	// потому что GORM-сервер использовал 8080.
	log.Fatal(http.ListenAndServe(":8081", nil))
}
