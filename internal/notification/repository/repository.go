package repository

import (
	"context"
	"time"

	"github.com/soorajsomans/notification-service/internal/notification/model"
)

type NotificationRepository interface {
	Create(
		ctx context.Context,
		notification *model.Notification,
	) error

	GetByID(
		ctx context.Context,
		id string,
	) (*model.Notification, error)

	FindPending(
		ctx context.Context,
		limit int,
	) ([]model.Notification, error)

	UpdateStatus(
		ctx context.Context,
		id string,
		status model.NotificationStatus,
	) error

	ClaimPendingNotifications(
		ctx context.Context,
		limit int,
	) ([]model.Notification, error)

	MarkForRetry(
		ctx context.Context,
		id string,
		retryCount int,
		nextRetryAt time.Time,
	) error
}
