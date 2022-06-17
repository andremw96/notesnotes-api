package api

import (
	db "andre/notesnotes-api/db/sqlc"
	"andre/notesnotes-api/util"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	Username  string `json:"username" binding:"required,alphanum"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
}

type createUserResponse struct {
	FullName   string         `json:"full_name"`
	FirstName  string         `json:"first_name"`
	LastName   sql.NullString `json:"last_name"`
	Username   string         `json:"username"`
	Email      string         `json:"email"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	NotesCount int32          `json:"notes_count"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUsersParams{
		FullName:  req.FirstName + " " + req.LastName,
		FirstName: req.FirstName,
		LastName:  sql.NullString{String: req.LastName, Valid: true},
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
	}

	user, err := server.store.CreateUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := createUserResponse{
		FullName:   user.FullName,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Username:   user.Username,
		Email:      user.Email,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		NotesCount: user.NotesCount,
	}

	ctx.JSON(http.StatusOK, response)
}

type getUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := createUserResponse{
		FullName:   user.FullName,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Username:   user.Username,
		Email:      user.Email,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		NotesCount: user.NotesCount,
	}

	ctx.JSON(http.StatusOK, response)
}

type listUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listUser(ctx *gin.Context) {
	var req listUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// offset is number of records database should skip
	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := server.store.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	responses := []createUserResponse{}
	for _, user := range users {
		response := createUserResponse{
			FullName:   user.FullName,
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			Username:   user.Username,
			Email:      user.Email,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
			NotesCount: user.NotesCount,
		}
		responses = append(responses, response)
	}

	ctx.JSON(http.StatusOK, responses)
}