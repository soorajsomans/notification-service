package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/soorajsomans/notification-service/internal/notification/model"
	"github.com/soorajsomans/notification-service/internal/notification/repository"
)

type NotificationServiceImpl struct {
	repo repository.NotificationRepository
}

func NewNotificationServiceImpl(repo repository.NotificationRepository) *NotificationServiceImpl {
	return &NotificationServiceImpl{
		repo: repo,
	}
}

func (s *NotificationServiceImpl) Create(
	ctx context.Context,
	n *model.Notification,
) error {

	if !n.Channel.IsValid() {
		return ErrInvalidChannel
	}
	now := time.Now()
	n.ID = uuid.NewString()
	n.Status = model.Pending
	n.CreatedAt = now
	n.UpdatedAt = now

	return s.repo.Create(ctx, n)
}

func (s *NotificationServiceImpl) GetByID(
	ctx context.Context,
	id string,
) (*model.Notification, error) {

	n, err := s.repo.GetByID(ctx, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotificationNotFound
		}

		return nil, err
	}
	return n, nil
}

var _ NotificationService = (*NotificationServiceImpl)(nil)
