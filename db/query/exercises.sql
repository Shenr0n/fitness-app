-- name: CreateExercise :one
INSERT INTO exercises (
    username,
    exercise_name,
    muscle_group
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetExercises :many
SELECT * FROM exercises
WHERE username = $1 
ORDER BY exer_id 
LIMIT $2 
OFFSET $3;

-- name: DeleteExercise :exec
DELETE FROM exercises
WHERE username = $1;