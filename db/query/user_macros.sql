-- name: RecordMacros :one
INSERT into user_macros (
    username,
    calories,
    fats,
    protein,
    carbs,
    um_date
) VALUES (
    $1, $2, $3, $4, $5, $6
) 
RETURNING *;

-- name: GetMacros :many
SELECT * FROM user_macros 
WHERE username = $1 
ORDER BY um_id 
LIMIT $2 
OFFSET $3;

-- name: GetMacroByDate :many
SELECT * FROM user_macros 
WHERE username = $1
	AND um_date = $2
LIMIT $3 
OFFSET $4;

-- name: DeleteUserMacros :exec
DELETE FROM user_macros
WHERE username = $1;