package service

import (
	"context"

	"github.com/soorajsomans/notification-service/internal/notification/model"
)

type NotificationService interface {
	Create(
		ctx context.Context,
		n *model.Notification,
	) error

	GetByID(
		ctx context.Context,
		id string,
	) (*model.Notification, error)
}
