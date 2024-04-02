package api

import (
	"errors"
	"net/http"

	db "github.com/Shenr0n/fitness-app/db/sqlc"
	"github.com/Shenr0n/fitness-app/token"
	"github.com/gin-gonic/gin"
)

func (server *Server) addExerciseToWorkout(ctx *gin.Context) {
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
	var reqExerToW addExerciseToWorkoutRequest
	if err := ctx.ShouldBindJSON(&reqExerToW); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	exer, errExer := server.store.GetExercise(ctx, reqExerToW.ExerID)
	workout, errWork := server.store.GetWorkout(ctx, reqExerToW.WorkoutID)

	if errExer != nil || errWork != nil {
		if errExer != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("Exercise not found")))
		} else {
			ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("Workout not found")))
		}
		return
	}
	if authPayload.Username != exer.Username || authPayload.Username != workout.Username || exer.Username != workout.Username {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("Username Mismatch: Unauthorized to make changes")))
		return
	}
	arg := db.AddExerciseToWorkoutParams{
		Username:  authPayload.Username,
		WorkoutID: reqExerToW.WorkoutID,
		ExerID:    reqExerToW.ExerID,
		Weights:   reqExerToW.Weights,
		Sets:      reqExerToW.Sets,
		Reps:      reqExerToW.Reps,
	}

	exerToW, err := server.store.AddExerciseToWorkout(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, exerToW)

}

func (server *Server) deleteExerciseInWE(ctx *gin.Context) {
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
	var reqExerInWorkout deleteExerciseInWorkoutRequest
	if err := ctx.ShouldBindJSON(&reqExerInWorkout); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	workout, errWork := server.store.GetWorkout(ctx, reqExerInWorkout.WorkoutID)
	exer, errExer := server.store.GetExercise(ctx, reqExerInWorkout.ExerID)

	if errWork != nil || errExer != nil {
		if errWork != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("Workout not found")))
		} else {
			ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("Exercise not found")))
		}
		return
	}
	if authPayload.Username != exer.Username || authPayload.Username != workout.Username || exer.Username != workout.Username {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("Username Mismatch: Unauthorized to make changes")))
		return
	}
	arg := db.DeleteExerciseInWorkoutWEParams{
		Username:  authPayload.Username,
		WorkoutID: reqExerInWorkout.WorkoutID,
		ExerID:    reqExerInWorkout.ExerID,
	}
	err := server.store.DeleteExerciseInWorkoutWE(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
