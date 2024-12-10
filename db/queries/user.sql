-- Check if the email already exists
-- name: CheckEmailExists :one
SELECT COUNT(*) 
FROM users 
WHERE email = $1;

-- Get user by email (used for login)
-- name: GetUserByEmail :one
SELECT user_id, username, email, password, created_at
FROM users
WHERE email = $1;

-- Create a new user
-- name: CreateUser :one
INSERT INTO users (username, email, password)
VALUES ($1, $2, $3)
RETURNING user_id, username, email, created_at;

-- Update user password
-- name: UpdateUserPassword :one
UPDATE users
SET password = $1, updated_at = $2
WHERE user_id = $3
RETURNING user_id, username, email, created_at, updated_at;

-- Get all users (for admin purposes)
-- name: GetAllUsers :many
SELECT user_id, username, email, created_at, updated_at
FROM users;
