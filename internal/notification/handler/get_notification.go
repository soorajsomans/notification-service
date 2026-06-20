package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetNotification(
	c *gin.Context,
) {
	id := c.Param("id")

	notification, err := h.service.GetByID(
		c.Request.Context(),
		id,
	)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		notification,
	)
}
