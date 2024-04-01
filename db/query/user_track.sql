-- name: RecordUserTrack :one
INSERT into user_track (
    username,
    weight,
    ut_date
) VALUES (
    $1, $2, $3
) 
RETURNING *;

-- name: GetUserTrack :many
SELECT ut_date, weight FROM user_track 
WHERE username = $1 
ORDER BY ut_date
LIMIT $2 
OFFSET $3;

-- name: DeleteUserTrack :exec
DELETE FROM user_track
WHERE username = $1;