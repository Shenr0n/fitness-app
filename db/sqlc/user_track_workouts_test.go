package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/Shenr0n/fitness-app/util"
	"github.com/stretchr/testify/require"
)

func randomUserTrackWorkout(t *testing.T) UserTrackWorkout {
	workout := createRandomWorkout(t)
	arg := RecordWorkoutParams{
		Username:    workout.Username,
		WorkoutID:   workout.WorkoutID,
		WorkoutName: workout.WorkoutName,
		UtwDate:     util.RandomDate(),
	}

	userTrackWorkout, err := testQueries.RecordWorkout(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userTrackWorkout)
	require.Equal(t, arg.Username, userTrackWorkout.Username)
	require.Equal(t, arg.WorkoutID, userTrackWorkout.WorkoutID)

	return userTrackWorkout
}

func TestRecordWorkout(t *testing.T) {
	randomUserTrackWorkout(t)
}

func TestGetRecords(t *testing.T) {
	record := randomUserTrackWorkout(t)
	require.NotEmpty(t, record)

	arg := GetRecordsParams{
		Username: record.Username,
		Limit:    10,
		Offset:   0,
	}
	userRecords, err := testQueries.GetRecords(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userRecords)

	for _, uR := range userRecords {
		require.NotEmpty(t, uR)
		fmt.Println("User records ", uR.WorkoutName, " ", uR.UtwDate)
	}

}
