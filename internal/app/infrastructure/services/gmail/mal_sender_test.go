package gmail

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMailSenderSendMail(t *testing.T) {
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	smtpEmail := os.Getenv("SMTP_EMAIL")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	to := os.Getenv("SMTP_TO_EMAIL")

	if smtpEmail == "" || smtpPassword == "" || to == "" {
		t.Skip("SMTP_EMAIL, SMTP_PASSWORD and SMTP_TO_EMAIL are required to run this test")
		return
	}

	sender := NewMailSender(smtpHost, smtpPort, smtpEmail, smtpPassword)

	subject := "Test"
	body := "Test of body"

	err := sender.SendMail(to, subject, body)

	assert.NoError(t, err)
}
