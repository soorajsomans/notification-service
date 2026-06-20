package dto

import "github.com/soorajsomans/notification-service/internal/notification/model"

type CreateNotificationRequest struct {
	UserId  string        `json:"userId" binding:"required"`
	Channel model.Channel `json:"channel" binding:"required"`
	Message string        `json:"message" binding:"required"`
}
