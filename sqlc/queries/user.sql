-- Создать пользователя
-- name: CreateUser :one
INSERT INTO users (name, email, password)
VALUES ($1, $2, $3)
RETURNING id, name, email, password;

-- Получить одного пользователя
-- name: GetUser :one
SELECT id, name, email, password
FROM users
WHERE id = $1;

-- Получить всех пользователей
-- name: GetUsers :many
SELECT id, name, email, password
FROM users
ORDER BY id;

-- Обновить пользователя
-- name: UpdateUser :one
UPDATE users
SET name = $2,
    email = $3,
    password = $4
WHERE id = $1
RETURNING id, name, email, password;

-- Удалить пользователя
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;