package handler

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(
	router *gin.RouterGroup,
) {
	router.POST(
		"/notifications",
		h.CreateNotification,
	)

	router.GET(
		"/notifications/:id",
		h.GetNotification,
	)

	router.GET(
		"/health",
		h.GetHealth,
	)
}
