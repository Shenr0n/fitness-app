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

-- name: DeleteUserTrackWorkouts :exec
DELETE FROM user_track_workouts
WHERE username = $1;