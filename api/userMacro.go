package api

import (
	"net/http"

	db "github.com/Shenr0n/fitness-app/db/sqlc"
	"github.com/Shenr0n/fitness-app/token"
	"github.com/gin-gonic/gin"
)

func (server *Server) recordMacros(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var reqMacro recordMacrosRequest
	if err := ctx.ShouldBindJSON(&reqMacro); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if req.Username != authPayload.Username {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}

	arg := db.RecordMacrosParams{
		Username: authPayload.Username,
		Calories: reqMacro.Calories,
		Fats:     reqMacro.Fats,
		Protein:  reqMacro.Protein,
		Carbs:    reqMacro.Carbs,
		UmDate:   reqMacro.UmDate,
	}

	macro, err := server.store.RecordMacros(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, macro)
}

func (server *Server) getMacros(ctx *gin.Context) {
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

	arg := db.GetMacrosParams{
		//Username: req.Username,
		Username: authPayload.Username,
		Limit:    reqPage.PageSize,
		Offset:   (reqPage.PageID - 1) * reqPage.PageSize,
	}
	macros, err := server.store.GetMacros(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	/*authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if req.Username != authPayload.Username {
		err := errors.New("You're unautorized to view other users")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}*/

	ctx.JSON(http.StatusOK, macros)
}

func (server *Server) getMacroByDate(ctx *gin.Context) {
	var req getUserAndDateRequest
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

	arg := db.GetMacroByDateParams{
		Username: authPayload.Username,
		UmDate:   req.UmDate,
		Limit:    reqPage.PageSize,
		Offset:   (reqPage.PageID - 1) * reqPage.PageSize,
	}
	macros, err := server.store.GetMacroByDate(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, macros)
}
