package di

import (
	"net/smtp"
	"os"
	"portal_back/mailer/api/internalapi"
	"portal_back/mailer/impl/app/mailsender"
	"portal_back/mailer/impl/app/notification"
	"portal_back/mailer/impl/app/passwordreseter"
)

func GetEnv(name, def string) string {
	s := os.Getenv(name)
	if s == "" {
		s = def
	}
	return s
}

func InitMailerModule(auth smtp.Auth) (internalapi.MailSenderService, internalapi.NotificationsService, internalapi.PasswordChangerService, error) {
	companyMail := GetEnv("COMPANY_MAIL", "teamtells@teamtells.ru")
	smtpPort := GetEnv("SMTP_PORT", "2525")
	addr := GetEnv("SMTP_SERVER", "localhost")

	mailSenderService := mailsender.NewService(auth, companyMail, smtpPort, addr)
	notificationsService, err := notification.NewService(mailSenderService)
	if err != nil {
		// TODO: wrap error
		return nil, nil, nil, err
	}
	passwordService, err := passwordreseter.NewService(mailSenderService)
	if err != nil {
		// TODO: wrap error
		return nil, nil, nil, err
	}
	return mailSenderService, notificationsService, passwordService, nil
}
