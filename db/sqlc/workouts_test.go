package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/Shenr0n/fitness-app/util"
	"github.com/stretchr/testify/require"
)

func createRandomWorkout(t *testing.T) Workout {
	user := createRandomUser(t)
	arg := CreateWorkoutParams{
		Username:    user.Username,
		WorkoutName: util.RandomText(),
	}
	workout, err := testQueries.CreateWorkout(context.Background(), arg)
	require.NotZero(t, workout.WorkoutID)
	require.NoError(t, err)
	require.NotEmpty(t, workout)
	fmt.Println("Workout user, id and name: ", workout.Username, ", ", workout.WorkoutID, ", ", workout.WorkoutName)
	return workout
}
func TestCreateWorkout(t *testing.T) {
	createRandomWorkout(t)
}

func TestGetWorkouts(t *testing.T) {
	workout := createRandomWorkout(t)
	arg := GetWorkoutsParams{
		Username: workout.Username,
		Limit:    5,
		Offset:   0,
	}
	workoutList, err := testQueries.GetWorkouts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, workoutList)
	fmt.Println("User: ", arg.Username)
	for _, workoutItem := range workoutList {
		require.NotEmpty(t, workoutItem)
		require.Equal(t, workout.Username, workoutItem.Username)
		fmt.Println("Workout id & name: ", workoutItem.WorkoutID, workoutItem.WorkoutName)
	}
}
