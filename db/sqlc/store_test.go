package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeleteExercise(t *testing.T) {

	store := NewStore(testDB)
	exerInWork := randomExerciseInWorkout(t)

	fmt.Println("User:", exerInWork.Username, " ExerID:", exerInWork.ExerID, " WorkoutID:", exerInWork.WorkoutID)
	errs := make(chan error)
	go func() {
		err := store.DeleteExercise(context.Background(), DeleteExerciseParams{
			Username: exerInWork.Username,
			ExerID:   exerInWork.ExerID,
		})

		errs <- err
	}()

	err := <-errs
	require.NoError(t, err)
}
func TestDeleteWorkout(t *testing.T) {
	store := NewStore(testDB)
	exerInWork := randomExerciseInWorkout(t)

	fmt.Println("User:", exerInWork.Username, " ExerID:", exerInWork.ExerID, " WorkoutID:", exerInWork.WorkoutID)
	errs := make(chan error)
	go func() {
		err := store.DeleteWorkout(context.Background(), DeleteWorkoutParams{
			Username:  exerInWork.Username,
			WorkoutID: exerInWork.WorkoutID,
		})

		errs <- err
	}()

	err := <-errs
	require.NoError(t, err)
}
func TestDeleteUser(t *testing.T) {
	store := NewStore(testDB)
	exerInWork := randomExerciseInWorkout(t)
	fmt.Println("User:", exerInWork.Username, " ExerID:", exerInWork.ExerID, " WorkoutID:", exerInWork.WorkoutID)
	errs := make(chan error)
	go func() {
		err := store.DeleteUser(context.Background(), exerInWork.Username)
		errs <- err
	}()
	err := <-errs
	require.NoError(t, err)
}
