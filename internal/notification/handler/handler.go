package handler

import (
	notificationService "github.com/soorajsomans/notification-service/internal/notification/service"
)

type Handler struct {
	service notificationService.NotificationService
}

func NewHandler(
	service notificationService.NotificationService,
) *Handler {
	return &Handler{
		service: service,
	}
}
