package gmail

import (
	"fmt"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/services"
	"github.com/jibaru/home-inventory-api/m/logger"
	"github.com/jibaru/home-inventory-api/m/notifier"
	"net/smtp"
)

type MailSender struct {
	smtpHost     string
	smtpPort     int
	smtpEmail    string
	smtpPassword string
}

func NewMailSender(
	smtpHost string,
	smtpPort int,
	smtpEmail string,
	smtpPassword string,
) *MailSender {
	return &MailSender{
		smtpHost:     smtpHost,
		smtpPort:     smtpPort,
		smtpEmail:    smtpEmail,
		smtpPassword: smtpPassword,
	}
}

func (ms *MailSender) hostPort() string {
	return fmt.Sprintf("%s:%d", ms.smtpHost, ms.smtpPort)
}

func (ms *MailSender) SendMail(to string, subject string, body string) error {
	from := ms.smtpEmail

	auth := smtp.PlainAuth("", from, ms.smtpPassword, ms.smtpHost)

	msg := []byte(
		"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" + body + "\r\n")

	err := smtp.SendMail(ms.hostPort(), auth, from, []string{to}, msg)
	if err != nil {
		logger.LogError(err)
		notifier.NotifyError(err)
		return services.ErrMailSenderCannotSendMail
	}

	return nil
}
