-- name: AddExerciseToWorkout :one
INSERT INTO workout_exercises (
    username,
    workout_id,
    exer_id,
    weights,
    sets,
    reps
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetWorkoutExercises :many
SELECT w.*, we.exer_id, e.exercise_name AS exercise_name, e.muscle_group AS muscle_group, we.weights, we.sets, we.reps
FROM workouts w
JOIN workout_exercises we ON w.workout_id = we.workout_id
JOIN exercises e ON we.exer_id = e.exer_id
WHERE w.username = $1
  AND w.workout_id = $2
ORDER BY we.exer_id
LIMIT $3 
OFFSET $4;

-- name: DeleteUserWorkoutExercises :exec
DELETE FROM workout_exercises
WHERE username = $1;

-- name: DeleteExerciseInWE :exec
DELETE FROM workout_exercises
WHERE username = $1
  AND exer_id = $2;

-- name: DeleteExerciseInWorkoutWE :exec
DELETE FROM workout_exercises
WHERE username = $1
  AND workout_id = $2
  AND exer_id = $3;

-- name: DeleteWorkoutInWE :exec
DELETE FROM workout_exercises
WHERE username = $1
  AND workout_id = $2;