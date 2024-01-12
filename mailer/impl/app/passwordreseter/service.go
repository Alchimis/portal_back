package passwordreseter

import (
	"context"
	"portal_back/mailer/impl/app/mailsender"
	"text/template"
)

type Service interface {
	SendPasswordResetMessage(ctx context.Context, email string, link string) error
	SendPasswordChangeMessage(ctx context.Context, email string, link string) error
}

type service struct {
	passwordResetTemplate  *template.Template
	passwordChangeTemplate *template.Template
	ms                     mailsender.Service
}

func NewService(r mailsender.Service) (Service, error) {
	passwordResetMessage, err := template.New("Send password reset template").Parse(`Follow the link to reset your password {{.Link}}`)
	if err != nil {
		// TODO: wrap error
		return nil, err
	}

	passwordChangeMessage, err := template.New("Send password change template").Parse(`Follow the link to change password {{.Link}}`)
	if err != nil {
		// TODO: wrapp error
		return nil, err
	}

	return &service{passwordResetMessage, passwordChangeMessage, r}, nil
}
func (s *service) SendPasswordResetMessage(ctx context.Context, email string, link string) error {
	return s.ms.SendMailFromCompany(ctx, "Reseting password", email, s.passwordResetTemplate, struct{ Link string }{Link: link})
}

func (s *service) SendPasswordChangeMessage(ctx context.Context, email string, link string) error {
	return s.ms.SendMailFromCompany(ctx, "Reseting password", email, s.passwordChangeTemplate, struct{ Link string }{Link: link})
}
