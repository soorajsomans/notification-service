package provider

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/soorajsomans/notification-service/internal/notification/model"
)

type EmailProvider struct{}

func NewEmailProvider() *EmailProvider {
	return &EmailProvider{}
}

func (p *EmailProvider) Send(
	ctx context.Context,
	notification model.Notification,
) error {

	if rand.Intn(2) == 0 {
		return errors.New("email provider timeout")
	}
	fmt.Printf(
		"\n EMAIL SENT \n"+
			"User: %s\n"+
			"Message: %s\n\n",
		notification.UserID,
		notification.Message,
	)
	return nil
}
