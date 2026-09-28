User CRUD

Учебный CRUD-сервис для регистрации и управления пользователями на Go.

О проекте

Проект представляет собой HTTP API для работы с пользователями.

Реализованы основные CRUD-операции:

* создание пользователя;
* получение пользователя по ID;
* получение списка пользователей;
* обновление пользователя;
* удаление пользователя.

Для работы с PostgreSQL реализованы три варианта:

1. GORM;
2. sqlx;
3. sqlc.

Также реализован собственный cache в памяти приложения и рассмотрены основные стратегии кеширования.

Технологии

* Go 1.27
* PostgreSQL
* Docker
* GORM
* sqlx
* sqlc
* net/http
* sync.RWMutex

Архитектура

Проект разделён на несколько уровней:

HTTP Handler
↓
Service
↓
Repository
↓
PostgreSQL

Handler

Принимает HTTP-запросы и формирует HTTP-ответы.

Service

Содержит основную логику приложения.

Repository

Отвечает за работу с базой данных.

Model

Описывает структуру пользователя.

Cache

Хранит данные в памяти приложения для быстрого получения.

Структура проекта

user-crud/
├── cmd/
│   └── main.go
│
├── internal/
│   ├── cache/
│   │   └── cache.go
│   ├── handler/
│   │   └── user_handler.go
│   ├── model/
│   │   └── user.go
│   ├── repository/
│   │   └── user_repository.go
│   └── service/
│       ├── user_service.go
│       └── cached_user_service.go
│
├── sqlx/
│   ├── handler/
│   │   └── user_handler.go
│   ├── model/
│   │   └── user.go
│   ├── repository/
│   │   └── user_repository.go
│   ├── service/
│   │   └── user_service.go
│   └── main.go
│
├── sqlc/
│   ├── generated/
│   │   ├── db.go
│   │   ├── models.go
│   │   └── user.sql.go
│   ├── handler/
│   │   └── user_handler.go
│   ├── queries/
│   │   └── user.sql
│   ├── repository/
│   │   └── user_repository.go
│   ├── service/
│   │   └── user_service.go
│   └── main.go
│
├── sqlc.yaml
├── go.mod
└── go.sum

CRUD API

Используются следующие HTTP-методы:

Метод	Endpoint	Назначение
POST	/users	Создание пользователя
GET	/users	Получение всех пользователей
GET	/users/{id}	Получение пользователя по ID
PUT	/users/{id}	Обновление пользователя
DELETE	/users/{id}	Удаление пользователя

Модель пользователя

Пользователь содержит следующие поля:

* ID
* Name
* Email
* Password

GORM

GORM используется как ORM для работы с PostgreSQL.

GORM позволяет работать с таблицами базы данных через Go-структуры и методы библиотеки.

Для создания и обновления таблицы используется:

db.AutoMigrate(&model.User{})

Запуск:

go run ./cmd

sqlx

sqlx используется для работы с PostgreSQL через обычные SQL-запросы.

Пример SQL-запроса:

SELECT id, name, email, password
FROM users
WHERE id = $1;

Запуск:

go run ./sqlx

sqlc

sqlc используется для генерации Go-кода на основе SQL-запросов.

SQL-запросы находятся в:

sqlc/queries/user.sql

Для генерации Go-кода используется команда:

sqlc generate

Сгенерированный код находится в:

sqlc/generated/

Запуск:

go run ./sqlc

Cache

В проекте реализован собственный cache в памяти приложения.

Для хранения данных используется:

map[uint]interface{}

Для безопасной работы с cache из нескольких goroutine используется:

sync.RWMutex

Cache поддерживает операции:

* Set — сохранить значение;
* Get — получить значение;
* Delete — удалить значение.

Стратегия Cache-Aside

В проекте используется стратегия Cache-Aside.

При получении пользователя сначала проверяется cache.

Запрос
↓
Cache
↓
Есть данные?
↙       ↘
Да       Нет
↓         ↓
Ответ   PostgreSQL
↓
Cache
↓
Ответ

Если пользователя нет в cache, приложение получает его из PostgreSQL и сохраняет в cache.

При обновлении пользователя данные обновляются в PostgreSQL и cache.

При удалении пользователя запись удаляется из PostgreSQL и cache.

Другие стратегии кеширования

В рамках проекта рассмотрены основные стратегии кеширования.

Cache-Aside

Приложение самостоятельно проверяет cache.

Если данных нет, приложение получает их из базы данных и сохраняет в cache.

Write-Through

При записи данные одновременно записываются в cache и базу данных.

Write-Back

Сначала данные записываются в cache, а затем сохраняются в базу данных.

Read-Through

Приложение обращается к cache, а cache самостоятельно получает отсутствующие данные из базы данных.

PostgreSQL

PostgreSQL используется как основная база данных проекта.

Название базы данных:

user_crud

PostgreSQL запускается в Docker.

Проверка проекта

Форматирование кода:

gofmt -w cmd internal sqlx sqlc

Генерация кода sqlc:

sqlc generate

Проверка сборки проекта:

go build ./…

Запуск

GORM:

go run ./cmd

sqlx:

go run ./sqlx

sqlc:

go run ./sqlc

Результат

В проекте реализован CRUD-сервис пользователей на Go с тремя вариантами работы с PostgreSQL:

* GORM;
* sqlx;
* sqlc.

Дополнительно реализован собственный потокобезопасный cache с использованием map и sync.RWMutex.

Проект предназначен для учебных целей и демонстрирует различные подходы к работе Go-приложения с PostgreSQL и кешированием.