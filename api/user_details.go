package api

import (
	"errors"
	"net/http"
	"time"

	db "github.com/Shenr0n/fitness-app/db/sqlc"
	"github.com/Shenr0n/fitness-app/token"
	"github.com/gin-gonic/gin"
)

func (server *Server) recordDetails(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if req.Username != authPayload.Username {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}
	var reqDetails recordDetailsRequest
	if err := ctx.ShouldBindJSON(&reqDetails); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	userTemp, err := server.store.GetUser(ctx, authPayload.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("Couldn't get user details internally")))
		return
	}
	ageTemp, err := calculateAge(userTemp.Dob)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("Couldn't calculate age")))
		return
	}

	details := db.RecordDetailsParams{
		Username:           authPayload.Username,
		Age:                int32(ageTemp),
		Weight:             userTemp.Weight,
		Height:             userTemp.Height,
		GoalWeight:         reqDetails.GoalWeight,
		DietPref:           reqDetails.DietPref,
		FoodAllergies:      reqDetails.FoodAllergies,
		DailyCalIntakeGoal: reqDetails.DailyCalIntakeGoal,
		ActivityLevel:      reqDetails.ActivityLevel,
		CurrentFitness:     reqDetails.CurrentFitness,
		FitnessGoal:        reqDetails.FitnessGoal,
	}

	sentDetails, err := server.store.RecordDetails(ctx, details)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sentDetails)
}

func calculateAge(dobStr string) (int, error) {

	dob, err := time.Parse("2006-01-02", dobStr)
	if err != nil {
		return 0, err
	}
	currentDate := time.Now()
	age := currentDate.Year() - dob.Year()

	// Check for birthday in current year
	if currentDate.Month() < dob.Month() ||
		(currentDate.Month() == dob.Month() && currentDate.Day() < dob.Day()) {
		age--
	}
	return age, nil
}
