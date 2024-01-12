package internalapi

import "context"

type PasswordChangerService interface {
	SendPasswordResetMessage(ctx context.Context, email string, link string) error
	SendPasswordChangeMessage(ctx context.Context, email string, link string) error
}
