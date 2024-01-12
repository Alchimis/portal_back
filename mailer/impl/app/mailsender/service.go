package mailsender

import (
	"context"
	"net/smtp"
	"strings"
	"text/template"
)

type Service interface {
	SendMail(ctx context.Context, subject string, recipientEmail string, senderEmail string, template *template.Template, payload interface{}) error
	SendMailFromCompany(ctx context.Context, subject string, recipientEmail string, template *template.Template, payload interface{}) error
}

type service struct {
	companyMail       string
	smtpPort          string
	smtpServerAddress string
	netAddress        string
	auth              smtp.Auth
}

func NewService(auth smtp.Auth, companyMail string, smtpPort string, smtpServerAddress string) Service {
	return &service{
		auth:              auth,
		companyMail:       companyMail,
		smtpPort:          smtpPort,
		smtpServerAddress: smtpServerAddress,
		netAddress:        smtpServerAddress + ":" + smtpPort,
	}
}

func (s *service) SendMailFromCompany(ctx context.Context, subject string, recipientEmail string, template *template.Template, payload interface{}) error {
	return s.SendMail(ctx, subject, recipientEmail, s.companyMail, template, payload)
}

func (s *service) SendMail(ctx context.Context, subject string, recipientEmail string, senderEmail string, template *template.Template, payload interface{}) error {
	message, err := compile(template, payload)
	if err != nil {
		// ToDo wrap error
		return err
	}

	m := joinStrings("From: ", senderEmail, "\r\n",
		"To: ", recipientEmail, "\r\n", "Subject: ", subject, "\r\n\r\n", message)
	return smtp.SendMail(s.netAddress, s.auth, senderEmail, []string{recipientEmail}, []byte(m))
}

func joinStrings(strs ...string) string {
	var b strings.Builder
	for _, s := range strs {
		b.WriteString(s)
	}
	return b.String()
}

func compile(t *template.Template, payload interface{}) (string, error) {
	var b strings.Builder

	err := t.Execute(&b, payload)
	if err != nil {
		// TODO: wrap error
		return "", err
	}
	return b.String(), nil
}
