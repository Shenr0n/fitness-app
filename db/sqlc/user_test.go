package db

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/Shenr0n/fitness-app/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomText(),
		HashedPassword: "secret",
		FullName:       util.RandomText(),
		Email:          util.RandomEmail(),
		Dob:            util.RandomDate(),
		Weight:         int32(util.RandomWeight()),
		Height:         int32(util.RandomHeight()),
	}
	//user, err := .CreateUser(context.Background(), arg)
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Dob, user.Dob)
	require.Equal(t, arg.Weight, user.Weight)
	require.Equal(t, arg.Height, user.Height)

	//require.NotZero(t, user.Username)
	require.NotZero(t, user.CreatedAt)

	err = testQueries.AddDefaultExercises(context.Background(), arg.Username)
	if err != nil {
		log.Fatal(err)
	}
	return user
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Dob, user2.Dob)
	require.Equal(t, user1.Weight, user2.Weight)
	require.Equal(t, user1.Height, user2.Height)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateHeight(t *testing.T) {
	user1 := createRandomUser(t)
	fmt.Println(user1.Username, "old height:", user1.Height)
	arg := UpdateHeightParams{
		Username: user1.Username,
		Height:   int32(util.RandomHeight()),
	}
	user2, err := testQueries.UpdateHeight(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, user1.Username, user2.Username)
	require.NotEqual(t, user1.Height, user2.Height)
	fmt.Println(user2.Username, "new height:", user2.Height)
}

func TestUpdateWeight(t *testing.T) {
	user1 := createRandomUser(t)
	fmt.Println(user1.Username, "old weight:", user1.Weight)
	arg := UpdateWeightParams{
		Username: user1.Username,
		Weight:   int32(util.RandomWeight()),
	}
	user2, err := testQueries.UpdateWeight(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, user1.Username, user2.Username)
	require.NotEqual(t, user1.Weight, user2.Weight)
	fmt.Println(user2.Username, "new height:", user2.Height)
}
