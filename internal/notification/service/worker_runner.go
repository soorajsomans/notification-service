package service

import (
	"context"
	"log"
	"time"
)

func StartWorker(
	worker *Worker,
) {
	log.Println("notification worker started")
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		err := worker.ProcessPendingNotifications(
			context.Background(),
		)

		if err != nil {
			log.Printf(
				"worker error : %v",
				err,
			)
		}
	}
}
