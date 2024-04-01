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

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// Execute function within a db transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

func (store *Store) DeleteExercise(ctx context.Context, arg DeleteExerciseParams) error {

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		err = q.DeleteExerciseInWE(ctx, DeleteExerciseInWEParams{
			Username: arg.Username,
			ExerID:   arg.ExerID,
		})
		if err != nil {
			return err
		}
		err = q.DeleteExercise(ctx, DeleteExerciseParams{
			Username: arg.Username,
			ExerID:   arg.ExerID,
		})
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

func (store *Store) DeleteWorkout(ctx context.Context, arg DeleteWorkoutParams) error {

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		err = q.DeleteWorkoutInWE(ctx, DeleteWorkoutInWEParams{
			Username:  arg.Username,
			WorkoutID: arg.WorkoutID,
		})
		if err != nil {
			return err
		}
		err = q.DeleteWorkout(ctx, DeleteWorkoutParams{
			Username:  arg.Username,
			WorkoutID: arg.WorkoutID,
		})
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

func (store *Store) DeleteUser(ctx context.Context, username string) error {

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		err = q.DeleteUserTrack(ctx, username)
		if err != nil {
			return err
		}
		err = q.DeleteUserMacros(ctx, username)
		if err != nil {
			return err
		}
		err = q.DeleteUserTrackWorkouts(ctx, username)
		if err != nil {
			return err
		}
		err = q.DeleteUserWorkoutExercises(ctx, username)
		if err != nil {
			return err
		}
		err = q.DeleteUserWorkouts(ctx, username)
		if err != nil {
			return err
		}
		err = q.DeleteUserExercises(ctx, username)
		if err != nil {
			return err
		}
		err = q.DeleteUser(ctx, username)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}
