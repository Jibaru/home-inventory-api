package services

import "errors"

var (
	ErrMailSenderCannotSendMail = errors.New("mail sender cannot send mail")
)

type MailSender interface {
	SendMail(to string, subject string, body string) error
}
