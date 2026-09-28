package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"user-crud/sqlc/generated"
	"user-crud/sqlc/handler"
	"user-crud/sqlc/repository"
	"user-crud/sqlc/service"

	_ "github.com/lib/pq"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=user_crud port=5432 sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("ошибка подключения к PostgreSQL:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("PostgreSQL не отвечает:", err)
	}

	fmt.Println("PostgreSQL подключён через sqlc!")

	queries := generated.New(db)

	userRepository := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepository)
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

	fmt.Println("SQLC сервер запущен на http://localhost:8082")

	log.Fatal(http.ListenAndServe(":8082", nil))
}
