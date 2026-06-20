package model

import "time"

type NotificationStatus string

const (
	Pending    NotificationStatus = "PENDING"
	Processing NotificationStatus = "PROCESSING"
	Sent       NotificationStatus = "SENT"
	Failed     NotificationStatus = "FAILED"
)

type Channel string

const (
	Email Channel = "EMAIL"
	SMS   Channel = "SMS"
	Push  Channel = "PUSH"
)

type Notification struct {
	ID     string `db:"id"`
	UserID string `db:"user_id"`

	Channel Channel `db:"channel"`
	Message string  `db:"message"`

	Status NotificationStatus `db:"status"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
