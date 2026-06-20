package service

import (
	"context"

	"github.com/soorajsomans/notification-service/internal/notification/model"
	"github.com/soorajsomans/notification-service/internal/notification/repository"
	"github.com/soorajsomans/notification-service/internal/provider"
)

type Worker struct {
	repo     repository.NotificationRepository
	provider provider.NotificationProvider
}

func NewWorker(
	repo repository.NotificationRepository,
	provider provider.NotificationProvider,
) *Worker {
	return &Worker{
		repo:     repo,
		provider: provider,
	}
}

func (w *Worker) ProcessPendingNotifications(
	ctx context.Context,
) error {
	notifications, err := w.repo.ClaimPendingNotifications(
		ctx,
		100,
	)

	if err != nil {
		return err
	}

	for _, notification := range notifications {
		err = w.provider.Send(
			ctx,
			notification,
		)

		if err != nil {
			_ = w.repo.UpdateStatus(
				ctx,
				notification.ID,
				model.Failed,
			)
			continue
		}

		_ = w.repo.UpdateStatus(
			ctx,
			notification.ID,
			model.Sent,
		)
	}
	return nil
}
