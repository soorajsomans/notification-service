package service

import "time"

func NextRetryTime(
	retryCount int,
) time.Time {
	delay :=
		time.Duration(
			1<<retryCount,
		) * time.Second
	return time.Now().Add(delay)
}
