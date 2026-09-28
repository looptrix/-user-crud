package main

import (
	"fmt"
	"log"
	"net/http"

	"user-crud/internal/cache"
	"user-crud/internal/handler"
	"user-crud/internal/model"
	"user-crud/internal/repository"
	"user-crud/internal/service"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=user_crud port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("ошибка подключения к PostgreSQL:", err)
	}

	fmt.Println("PostgreSQL подключён!")

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("ошибка миграции:", err)
	}

	userRepository := repository.NewUserRepository(db)

	// Создаём cache в памяти приложения.
	userCache := cache.New()

	// Используем сервис с cache.
	userService := service.NewCachedUserService(
		userRepository,
		userCache,
	)

	userHandler := handler.NewUserHandler(userService)

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

	fmt.Println("GORM + Cache сервер запущен на http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
