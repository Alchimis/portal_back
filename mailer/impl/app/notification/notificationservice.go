package notification

import (
	"context"
	"portal_back/mailer/impl/app/mailsender"
	"text/template"
	"time"
)

type Service interface {
	SendSubscriptionExpirationReminder(ctx context.Context, email string, expirationDate time.Time) error
}

type service struct {
	subscriptionExpirationTemplate *template.Template
	ms                             mailsender.Service
}

func NewService(ms mailsender.Service) (Service, error) {

	t, err := template.New("SubscriptionExpiration").Parse("")
	if err != nil {
		// TODO: wrap error
		return nil, err
	}

	return &service{t, ms}, nil
}

func (s *service) SendSubscriptionExpirationReminder(ctx context.Context, email string, expirationDate time.Time) error {
	return s.ms.SendMailFromCompany(ctx, "Subscription expiration", email, s.subscriptionExpirationTemplate, nil)
}
