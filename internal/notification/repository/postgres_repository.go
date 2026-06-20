package repository

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/soorajsomans/notification-service/internal/notification/model"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(
	db *sqlx.DB,
) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(
	ctx context.Context,
	n *model.Notification,
) error {
	query := `
	INSERT INTO notifications
	(
		id,
		user_id,
		channel,
		message,
		status,
		created_at,
		updated_at
	)
	VALUES
	($1,$2,$3,$4,$5,$6,$7)
	`
	_, err := r.db.ExecContext(
		ctx,
		query,
		n.ID,
		n.UserID,
		n.Channel,
		n.Message,
		n.Status,
		n.CreatedAt,
		n.UpdatedAt,
	)
	return err
}

func (r *PostgresRepository) GetByID(
	ctx context.Context,
	id string,
) (*model.Notification, error) {
	query := `
	SELECT
		id,
		user_id,
		channel,
		message,
		status,
		created_at,
		updated_at
	FROM notifications
	WHERE id=$1
	`

	var n model.Notification

	err := r.db.GetContext(
		ctx,
		&n,
		query,
		id,
	)

	if err != nil {
		return nil, err
	}
	return &n, nil
}

func (r *PostgresRepository) FindPending(
	ctx context.Context,
	limit int,
) ([]model.Notification, error) {
	query := `
	SELECT
		id,
		user_id,
		channel,
		message,
		status,
		created_at,
		updated_at
	FROM notifications
	WHERE status = 'PENDING'
	FOR UPDATE SKIP LOCKED
	ORDER BY created_at
	LIMIT $1
	`
	var notifications []model.Notification

	err := r.db.SelectContext(
		ctx,
		&notifications,
		query,
		limit,
	)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (r *PostgresRepository) UpdateStatus(
	ctx context.Context,
	id string,
	status model.NotificationStatus,
) error {
	query := `
	UPDATE notifications
	SET
		status = $1,
		updated_at= NOW()
	WHERE id = $2
	`
	_, err := r.db.ExecContext(
		ctx,
		query,
		status,
		id,
	)
	return err

}

func (r *PostgresRepository) ClaimPendingNotifications(
	ctx context.Context,
	limit int,
) ([]model.Notification, error) {
	query := `
	WITH claimed AS(
		SELECT id
		FROM notifications
		WHERE 
		(
			status = 'PENDING'
		)
		OR
		(
			status = 'RETRY'
			AND next_retry_at <= NOW()
		)
		ORDER BY created_at
		FOR UPDATE SKIP LOCKED
		LIMIT $1
	)
	UPDATE notifications n
	SET
		status='PROCESSING',
		updated_at = NOW()
	FROM claimed
	WHERE n.id = claimed.id
	RETURNING
		n.id,
		n.user_id,
		n.channel,
		n.message,
		n.status,
		n.created_at,
		n.updated_at,
		n.retry_count,
		n.next_retry_at
	`

	var notifications []model.Notification

	err := r.db.SelectContext(
		ctx,
		&notifications,
		query,
		limit,
	)

	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (r *PostgresRepository) MarkForRetry(
	ctx context.Context,
	id string,
	retryCount int,
	nextRetryAt time.Time,
) error {
	query := `
	UPDATE notifications
	SET
		status = 'RETRY',
		retry_count = $1,
		next_retry_at = $2,
		updated_at = NOW()
	WHERE id = $3
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		retryCount,
		nextRetryAt,
		id,
	)
	rows, _ := result.RowsAffected()
	log.Printf("rows updated = %d", rows)
	return err
}

var _ NotificationRepository = (*PostgresRepository)(nil)
