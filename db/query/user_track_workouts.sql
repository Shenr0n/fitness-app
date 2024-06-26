-- name: RecordWorkout :one
INSERT into user_track_workouts (
    username,
    workout_id,
    workout_name,
    utw_date
) VALUES (
    $1, $2, $3, $4
) 
RETURNING *;

-- name: GetRecords :many
SELECT utw_id, workout_id, workout_name, utw_date 
FROM user_track_workouts
WHERE username = $1
ORDER BY utw_date
LIMIT $2
OFFSET $3;

-- name: GetRecord :one
SELECT username
FROM user_track_workouts
WHERE utw_id = $1;

-- name: DeleteUserWorkoutRecord :exec
DELETE FROM user_track_workouts
WHERE utw_id = $1;

-- name: DeleteUserTrackWorkouts :exec
DELETE FROM user_track_workouts
WHERE username = $1;