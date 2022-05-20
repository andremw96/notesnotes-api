package handler

import (
	"andre/notesnotes-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RemoveBookHandler(c *gin.Context) {
	id := c.Param("id")

	if result := h.db.Delete(&model.Book{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
