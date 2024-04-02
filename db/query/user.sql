-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  full_name,
  email,
  dob,
  weight,
  height
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: AddDefaultExercises :exec
INSERT INTO exercises (username, exercise_name, muscle_group, created_at)
SELECT $1, exercise_name, muscle_group, NOW()
FROM default_exercises;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateWeight :one
UPDATE users
SET 
weight = $2
WHERE username = $1
RETURNING *;

-- name: UpdateHeight :one
UPDATE users
SET 
height = $2
WHERE username = $1
RETURNING *;

-- name: UpdatePassword :exec
UPDATE users
SET 
hashed_password = $2
WHERE username = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;