package api

import (
	db "andre/notesnotes-api/db/sqlc"
	"andre/notesnotes-api/token"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type insertNewNoteRequest struct {
	UserID      int32  `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

func (server *Server) insertNewNote(ctx *gin.Context) {
	var req insertNewNoteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	loggedInUser, valid := server.validUser(ctx, req.UserID)
	if !valid {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if loggedInUser.Username != authPayload.Username {
		err := errors.New("Logged in user is different with authorization bearer")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.InsertNoteTxParams{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
	}

	result, err := server.store.InsertNewNote(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validUser(ctx *gin.Context, userID int32) (*db.User, bool) {
	user, err := server.store.GetUser(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return nil, false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return nil, false
	}

	return &user, true
}
