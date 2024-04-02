package api

import (
	"net/http"

	db "github.com/Shenr0n/fitness-app/db/sqlc"
	"github.com/Shenr0n/fitness-app/token"
	"github.com/gin-gonic/gin"
)

func (server *Server) createWorkout(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var reqWorkout createWorkoutRequest
	if err := ctx.ShouldBindJSON(&reqWorkout); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if req.Username != authPayload.Username {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}

	arg := db.CreateWorkoutParams{
		Username:    authPayload.Username,
		WorkoutName: reqWorkout.WorkoutName,
	}

	workout, err := server.store.CreateWorkout(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, workout)
}

func (server *Server) getWorkouts(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var reqPage getPageRequest
	if err := ctx.ShouldBindQuery(&reqPage); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if req.Username != authPayload.Username {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}
	arg := db.GetWorkoutsParams{
		Username: authPayload.Username,
		Limit:    reqPage.PageSize,
		Offset:   (reqPage.PageID - 1) * reqPage.PageSize,
	}

	workouts, err := server.store.GetWorkouts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, workouts)
}

func (server *Server) deleteWorkout(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var reqWorkoutID deleteWorkoutRequest
	if err := ctx.ShouldBindJSON(&reqWorkoutID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if req.Username != authPayload.Username {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}
	arg := db.DeleteWorkoutParams{
		Username:  authPayload.Username,
		WorkoutID: reqWorkoutID.WorkoutID,
	}
	err := server.store.DeleteWorkout(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
