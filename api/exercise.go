package api

import (
	"net/http"

	db "github.com/Shenr0n/fitness-app/db/sqlc"
	"github.com/Shenr0n/fitness-app/token"
	"github.com/gin-gonic/gin"
)

func (server *Server) createExercise(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var reqExer createExerciseRequest
	if err := ctx.ShouldBindJSON(&reqExer); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if req.Username != authPayload.Username {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}
	arg := db.CreateExerciseParams{
		Username:     authPayload.Username,
		ExerciseName: reqExer.ExerciseName,
		MuscleGroup:  reqExer.MuscleGroup,
	}

	exer, err := server.store.CreateExercise(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, exer)
}

func (server *Server) getExercises(ctx *gin.Context) {
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

	arg := db.GetExercisesParams{
		Username: authPayload.Username,
		Limit:    reqPage.PageSize,
		Offset:   (reqPage.PageID - 1) * reqPage.PageSize,
	}

	exer, err := server.store.GetExercises(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, exer)
}

func (server *Server) deleteExercise(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var reqExerID getExerciseRequest
	if err := ctx.ShouldBindJSON(&reqExerID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if req.Username != authPayload.Username {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}
	arg := db.DeleteExerciseParams{
		Username: authPayload.Username,
		ExerID:   reqExerID.ExerID,
	}
	err := server.store.DeleteExercise(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
