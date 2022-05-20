package handler

import (
	"andre/notesnotes-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ListBooksHandler(c *gin.Context) {
	var books []model.Book

	if result := h.db.Find(&books); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &books)

}
