// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user_track.sql

package db

import (
	"context"
)

const deleteUserTrack = `-- name: DeleteUserTrack :exec
DELETE FROM user_track
WHERE username = $1
`

func (q *Queries) DeleteUserTrack(ctx context.Context, username string) error {
	_, err := q.db.ExecContext(ctx, deleteUserTrack, username)
	return err
}

const getUserTrack = `-- name: GetUserTrack :many
SELECT ut_id, username, weight, ut_date, created_at FROM user_track 
WHERE username = $1 
ORDER BY ut_id 
LIMIT $2 
OFFSET $3
`

type GetUserTrackParams struct {
	Username string `json:"username"`
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
}

func (q *Queries) GetUserTrack(ctx context.Context, arg GetUserTrackParams) ([]UserTrack, error) {
	rows, err := q.db.QueryContext(ctx, getUserTrack, arg.Username, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserTrack{}
	for rows.Next() {
		var i UserTrack
		if err := rows.Scan(
			&i.UtID,
			&i.Username,
			&i.Weight,
			&i.UtDate,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const recordUserTrack = `-- name: RecordUserTrack :one
INSERT into user_track (
    username,
    weight,
    ut_date
) VALUES (
    $1, $2, $3
) 
RETURNING ut_id, username, weight, ut_date, created_at
`

type RecordUserTrackParams struct {
	Username string `json:"username"`
	Weight   int32  `json:"weight"`
	UtDate   string `json:"ut_date"`
}

func (q *Queries) RecordUserTrack(ctx context.Context, arg RecordUserTrackParams) (UserTrack, error) {
	row := q.db.QueryRowContext(ctx, recordUserTrack, arg.Username, arg.Weight, arg.UtDate)
	var i UserTrack
	err := row.Scan(
		&i.UtID,
		&i.Username,
		&i.Weight,
		&i.UtDate,
		&i.CreatedAt,
	)
	return i, err
}
