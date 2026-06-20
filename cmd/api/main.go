package main

import (
	"log"

	"github.com/gin-gonic/gin"
	notificationHandler "github.com/soorajsomans/notification-service/internal/notification/handler"
	"github.com/soorajsomans/notification-service/internal/notification/repository"
	"github.com/soorajsomans/notification-service/internal/notification/service"
	notificationService "github.com/soorajsomans/notification-service/internal/notification/service"
	"github.com/soorajsomans/notification-service/internal/platform/database"
	"github.com/soorajsomans/notification-service/internal/provider"
)

func main() {
	db, err := database.NewPostgres()

	if err != nil {
		panic(err)
	}

	repo := repository.NewPostgresRepository(db)

	svc := notificationService.NewNotificationServiceImpl(repo)
	handler := notificationHandler.NewHandler(svc)

	router := gin.Default()

	api := router.Group("/api/v1")

	handler.RegisterRoutes(api)

	emailProvider := provider.NewEmailProvider()

	worker := service.NewWorker(
		repo,
		emailProvider,
	)

	//start worker
	go service.StartWorker(
		worker,
	)

	log.Println("server started on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
