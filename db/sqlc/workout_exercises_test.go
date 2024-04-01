package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/Shenr0n/fitness-app/util"
	"github.com/stretchr/testify/require"
)

func randomExerciseInWorkout(t *testing.T) WorkoutExercise {
	exer := createRandomExercise(t)
	require.NotZero(t, exer.ExerID)
	require.NotEmpty(t, exer)
	require.NotZero(t, exer.CreatedAt)
	argWorkout := CreateWorkoutParams{
		Username:    exer.Username,
		WorkoutName: util.RandomText(),
	}
	workout, err := testQueries.CreateWorkout(context.Background(), argWorkout)
	require.NotZero(t, workout.WorkoutID)
	require.NoError(t, err)
	require.NotEmpty(t, workout)
	require.NotZero(t, workout.CreatedAt)
	require.Equal(t, workout.Username, exer.Username)
	//fmt.Println("Workout user, id and name: ", workout.Username, ", ", workout.WorkoutID, ", ", workout.WorkoutName)
	argWorkExer := AddExerciseToWorkoutParams{
		Username:  workout.Username,
		WorkoutID: workout.WorkoutID,
		ExerID:    exer.ExerID,
		Weights:   int32(util.RandomNum()),
		Sets:      int32(util.RandomNum()),
		Reps:      int32(util.RandomNum()),
	}
	workExer, err := testQueries.AddExerciseToWorkout(context.Background(), argWorkExer)
	require.NotZero(t, workExer.WeID)
	require.NoError(t, err)
	require.NotEmpty(t, workExer)
	require.Equal(t, workExer.Username, workout.Username)
	require.Equal(t, workExer.Username, exer.Username)
	require.NotZero(t, workExer.CreatedAt)
	return workExer
}

func TestAddExerciseToWorkout(t *testing.T) {
	randomExerciseInWorkout(t)
}

func TestGetWorkoutExercises(t *testing.T) {
	workExer := randomExerciseInWorkout(t)
	arg := GetWorkoutExercisesParams{
		Username:  workExer.Username,
		WorkoutID: 1,
		Limit:     10,
		Offset:    0,
	}
	workoutExerList, err := testQueries.GetWorkoutExercises(context.Background(), arg)
	require.NoError(t, err)
	//require.NotEmpty(t, workoutExerList)
	fmt.Println("User and workout ", arg.Username, arg.WorkoutID)
	for _, workoutExerItem := range workoutExerList {
		//require.NotEmpty(t, workoutExerItem)
		require.Equal(t, "vkbtms", workoutExerItem.Username)
		fmt.Println("Workout id, name: ", workoutExerItem.WorkoutName)
	}
}
