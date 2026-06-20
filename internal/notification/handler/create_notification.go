package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soorajsomans/notification-service/internal/notification/dto"
	"github.com/soorajsomans/notification-service/internal/notification/model"
)

func (h *Handler) CreateNotification(
	c *gin.Context,
) {

	var req dto.CreateNotificationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	notification := &model.Notification{
		UserID:  req.UserId,
		Channel: req.Channel,
		Message: req.Message,
	}
	err := h.service.Create(
		c.Request.Context(),
		notification,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		dto.CreateNotificationResponse{
			ID: notification.ID,
		},
	)
}
