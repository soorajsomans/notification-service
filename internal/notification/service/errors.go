package service

import "errors"

var (
	ErrInvalidChannel       = errors.New("invalid channel")
	ErrNotificationNotFound = errors.New("notification not found")
)
