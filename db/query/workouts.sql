-- name: CreateWorkout :one
INSERT INTO workouts (
  username,
  workout_name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetWorkouts :many
SELECT * FROM workouts
WHERE username = $1 
ORDER BY workout_id 
LIMIT $2 
OFFSET $3;

-- name: DeleteWorkout :exec
DELETE FROM workouts
WHERE username = $1
  AND workout_id = $2;

-- name: DeleteUserWorkouts :exec
DELETE FROM workouts
WHERE username = $1;