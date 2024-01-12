package internalapi

import (
	"context"
	"text/template"
)

type MailSenderService interface {
	SendMail(ctx context.Context, subject string, recipientEmail string, senderEmail string, template *template.Template, payload interface{}) error
	SendMailFromCompany(ctx context.Context, subject string, recipientEmail string, template *template.Template, payload interface{}) error
}
