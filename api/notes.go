package api

import (
	db "andre/notesnotes-api/db/sqlc"
	"database/sql"
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

	if !server.validUser(ctx, req.UserID) {
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

func (server *Server) validUser(ctx *gin.Context, userID int32) bool {
	_, err := server.store.GetUser(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	return true
}
