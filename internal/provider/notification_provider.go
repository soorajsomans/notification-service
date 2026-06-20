package provider

import (
	"context"

	"github.com/soorajsomans/notification-service/internal/notification/model"
)

type NotificationProvider interface {
	Send(
		ctx context.Context,
		notification model.Notification,
	) error
}
