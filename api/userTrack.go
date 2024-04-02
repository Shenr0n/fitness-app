package api

import (
	"net/http"

	db "github.com/Shenr0n/fitness-app/db/sqlc"
	"github.com/Shenr0n/fitness-app/token"
	"github.com/gin-gonic/gin"
)

func (server *Server) recordUserTrack(ctx *gin.Context) {
	var req getUserRequest
	var reqUserTrack recordUserTrackRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&reqUserTrack); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if req.Username != authPayload.Username {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}

	arg := db.RecordUserTrackParams{
		Username: authPayload.Username,
		Weight:   reqUserTrack.Weight,
		UtDate:   reqUserTrack.UtDate,
	}

	userTrack, err := server.store.RecordUserTrack(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, userTrack)
}

func (server *Server) getUserTrack(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var reqPage getPageRequest
	if err := ctx.ShouldBindQuery(&reqPage); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if req.Username != authPayload.Username {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}
	arg := db.GetUserTrackParams{
		Username: authPayload.Username,
		Limit:    reqPage.PageSize,
		Offset:   (reqPage.PageID - 1) * reqPage.PageSize,
	}
	userTrack, err := server.store.GetUserTrack(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userTrack)
}
