package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DeleteUserInfoParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

/*type Store interface {
	Querier
	DeleteUserInfo(ctx context.Context, arg DeleteUserInfoParams) error
}*/

func (store *SQLStore) DeleteUserInfo(ctx context.Context, arg DeleteUserInfoParams) error {
	// Begin a transaction
	tx, err := store.db.Begin()
	if err != nil {
		return err
	}
	// Defer transaction rollback or commit
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	// Execute delete operations within the transaction
	err = store.execTx(ctx, tx, func(q *Queries) error {

		// Delete user's exercises
		err := q.DeleteExercise(ctx, arg.Username)
		if err != nil {
			return err
		}

		// Delete user's workouts
		err = q.DeleteWorkouts(ctx, arg.Username)
		if err != nil {
			return err
		}

		// Delete user's workout exercises
		err = q.DeleteWorkoutExercises(ctx, arg.Username)
		if err != nil {
			return err
		}

		// Delete user's tracked workouts
		err = q.DeleteUserTrackWorkouts(ctx, arg.Username)
		if err != nil {
			return err
		}

		// Delete user's track info
		err = q.DeleteUserTrack(ctx, arg.Username)
		if err != nil {
			return err
		}

		// Delete user's macros
		err = q.DeleteUserMacros(ctx, arg.Username)
		if err != nil {
			return err
		}

		// Delete the user
		err = q.DeleteUser(ctx, arg.Username)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

// Execute function within a db transaction
func (store *SQLStore) execTx(ctx context.Context, tx *sql.Tx, fn func(*Queries) error) error {
	q := New(tx)
	err := fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
