package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/Shenr0n/fitness-app/util"
	"github.com/stretchr/testify/require"
)

func createRandomMacro(t *testing.T) UserMacro {
	user1 := createRandomUser(t)
	arg := RecordMacrosParams{
		Username: user1.Username,
		Calories: int32(util.RandomCal()),
		Fats:     int32(util.RandomQuantity()),
		Protein:  int32(util.RandomQuantity()),
		Carbs:    int32(util.RandomQuantity()),
		UmDate:   util.RandomDate(),
	}
	user1Macro, err := testQueries.RecordMacros(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user1Macro)

	require.Equal(t, arg.Username, user1Macro.Username)
	require.Equal(t, arg.Calories, user1Macro.Calories)
	require.Equal(t, arg.Fats, user1Macro.Fats)
	require.Equal(t, arg.Protein, user1Macro.Protein)
	require.Equal(t, arg.Carbs, user1Macro.Carbs)
	require.Equal(t, arg.UmDate, user1Macro.UmDate)

	require.NotZero(t, user1Macro.UmID)
	require.NotZero(t, user1Macro.CreatedAt)

	return user1Macro
}
func TestRecordMacros(t *testing.T) {
	createRandomMacro(t)
}

func TestGetMacros(t *testing.T) {
	userMacro := createRandomMacro(t)
	arg := GetMacrosParams{
		Username: userMacro.Username,
		Limit:    5,
		Offset:   0,
	}
	userMacros, err := testQueries.GetMacros(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userMacros)

	for _, uM := range userMacros {
		require.NotEmpty(t, uM)
		require.Equal(t, userMacro.Username, uM.Username)
		fmt.Println(uM.Username, " ", uM.UmDate, " ", uM.Calories, " ", uM.Carbs, " ", uM.Fats, " ", uM.Protein)
	}
}

func TestGetMacroByDate(t *testing.T) {
	userMacro := createRandomMacro(t)
	arg := GetMacroByDateParams{
		Username: userMacro.Username,
		UmDate:   userMacro.UmDate,
		Limit:    5,
		Offset:   0,
	}
	macroByDate, err := testQueries.GetMacroByDate(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, macroByDate)

	for _, macroDt := range macroByDate {
		require.NotEmpty(t, macroDt)
		require.Equal(t, userMacro.Username, macroDt.Username)
		require.Equal(t, userMacro.UmDate, macroDt.UmDate)
		fmt.Println(macroDt.Username, " ", macroDt.UmDate, " ", macroDt.Calories, " ", macroDt.Carbs, " ", macroDt.Fats, " ", macroDt.Protein)
	}
}
