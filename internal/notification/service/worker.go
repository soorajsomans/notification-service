package service

import (
	"context"
	"log"

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
		log.Printf(
			"claimed id=%s status=%s retry=%d nextRetry=%v",
			notification.ID,
			notification.Status,
			notification.RetryCount,
			notification.NextRetryAt,
		)
		err = w.provider.Send(
			ctx,
			notification,
		)

		if err != nil {
			log.Println("provider is down retrying")
			retryCount :=
				notification.RetryCount + 1

			nextRetryAt := NextRetryTime(retryCount)

			log.Printf(
				"Marking for retry. id=%s retryCount=%d nextRetry=%v",
				notification.ID,
				retryCount,
				nextRetryAt,
			)

			if retryCount >= MaxRetryCount {
				_ = w.repo.UpdateStatus(
					ctx,
					notification.ID,
					model.Failed,
				)
				continue
			}

			err = w.repo.MarkForRetry(
				ctx,
				notification.ID,
				retryCount,
				nextRetryAt,
			)

			if err != nil {
				log.Printf(
					"failed to mark retry: %v",
					err,
				)
			} else {
				log.Printf("Successfully marked for retry")
			}
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
