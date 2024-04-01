package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/Shenr0n/fitness-app/util"
	"github.com/stretchr/testify/require"
)

func createRandomExercise(t *testing.T) Exercise {
	user := createRandomUser(t)
	arg := CreateExerciseParams{

		Username:     user.Username,
		ExerciseName: util.RandomText(),
		MuscleGroup:  util.RandomText(),
	}
	exer, err := testQueries.CreateExercise(context.Background(), arg)
	require.NotZero(t, exer.ExerID)
	require.NoError(t, err)
	require.NotEmpty(t, exer)
	require.Equal(t, arg.Username, exer.Username)
	require.NotZero(t, exer.CreatedAt)
	return exer
}
func TestCreateExercise(t *testing.T) {
	createRandomExercise(t)
}

func TestGetExercises(t *testing.T) {
	exer := createRandomExercise(t)
	arg := GetExercisesParams{
		Username: exer.Username,
		Limit:    15,
		Offset:   0,
	}
	exerList, err := testQueries.GetExercises(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, exerList)
	fmt.Println("User: ", arg.Username)
	for _, exerItem := range exerList {
		require.NotEmpty(t, exerItem)
		require.Equal(t, arg.Username, exerItem.Username)
		fmt.Println("Exercise name: ", exerItem.ExerciseName)
	}
}

/*
func TestDeleteExercise(t *testing.T) {
	exer := createRandomExercise(t)
	require.NotEmpty(t, exer)
	fmt.Println("ID: ", exer.ExerID, " username: ", exer.Username)
	err := testQueries.DeleteExercise(context.Background(), exer.Username)
	require.NoError(t, err)
}*/
