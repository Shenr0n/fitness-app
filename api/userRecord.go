package api

import (
	"errors"
	"net/http"

	db "github.com/Shenr0n/fitness-app/db/sqlc"
	"github.com/Shenr0n/fitness-app/token"
	"github.com/gin-gonic/gin"
)

func newRecordResponse(rr db.GetRecordsRow) recordResponse {
	return recordResponse{
		UtwID:       rr.UtwID,
		WorkoutID:   rr.WorkoutID,
		WorkoutName: rr.WorkoutName,
		UtwDate:     rr.UtwDate,
	}
}
func (server *Server) recordUserWorkout(ctx *gin.Context) {
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

	var reqUserWorkout recordUserWorkoutRequest
	if err := ctx.ShouldBindJSON(&reqUserWorkout); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	workoutCheck, errWork := server.store.GetWorkout(ctx, reqUserWorkout.WorkoutID)
	if errWork != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errWork))
		return
	}
	if authPayload.Username != workoutCheck.Username {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("Workout not found")))
		return
	}

	arg := db.RecordWorkoutParams{
		Username:    authPayload.Username,
		WorkoutID:   reqUserWorkout.WorkoutID,
		WorkoutName: workoutCheck.WorkoutName,
		UtwDate:     reqUserWorkout.UtwDate,
	}

	recordWorkout, err := server.store.RecordWorkout(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, recordWorkout)

}

func (server *Server) deleteUserWorkoutRecord(ctx *gin.Context) {
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

	var reqUtw getUtwIDRequest
	if err := ctx.ShouldBindJSON(&reqUtw); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	usernameFromRecord, err := server.store.GetRecord(ctx, reqUtw.UtwID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("Record not found")))
		return
	}
	if authPayload.Username != usernameFromRecord {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("Unauthorized")))
		return
	}

	err = server.store.DeleteUserWorkoutRecord(ctx, reqUtw.UtwID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (server *Server) getUserWorkoutRecords(ctx *gin.Context) {
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
	var reqPage getPageRequest
	if err := ctx.ShouldBindQuery(&reqPage); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.GetRecordsParams{
		Username: authPayload.Username,
		Limit:    reqPage.PageSize,
		Offset:   (reqPage.PageID - 1) * reqPage.PageSize,
	}

	records, err := server.store.GetRecords(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	var rsp []recordResponse
	for _, record := range records {
		rsp = append(rsp, newRecordResponse(record))
	}

	ctx.JSON(http.StatusOK, rsp)
}
