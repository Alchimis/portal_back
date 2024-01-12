package internalapi

import (
	"context"
	"time"
)

type NotificationsService interface {
	SendSubscriptionExpirationReminder(ctx context.Context, email string, expirationDate time.Time) error
}
