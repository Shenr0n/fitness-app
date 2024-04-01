package db

import (
	"context"
	"testing"

	"github.com/Shenr0n/fitness-app/util"
	"github.com/stretchr/testify/require"
)

func randomUserTrack(t *testing.T) UserTrack {
	user := createRandomUser(t)
	arg := RecordUserTrackParams{
		Username: user.Username,
		Weight:   int32(util.RandomWeight()),
		UtDate:   util.RandomDate(),
	}
	userTrack, err := testQueries.RecordUserTrack(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userTrack)
	require.Equal(t, arg.Username, userTrack.Username)
	require.Equal(t, arg.Weight, userTrack.Weight)
	require.Equal(t, arg.UtDate, userTrack.UtDate)
	return userTrack
}

func TestRecordUserTrack(t *testing.T) {
	randomUserTrack(t)
}

func TestGetUserTrack(t *testing.T) {
	userTrack := randomUserTrack(t)
	arg := GetUserTrackParams{
		Username: userTrack.Username,
		Limit:    10,
		Offset:   0,
	}
	uTList, err := testQueries.GetUserTrack(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, uTList)
	require.Equal(t, userTrack.Username, arg.Username)
	for _, uT := range uTList {
		require.NotEmpty(t, uT)
	}
}
