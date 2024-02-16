package stub

import "github.com/stretchr/testify/mock"

type MailSenderMock struct {
	mock.Mock
}

func (m *MailSenderMock) SendMail(to string, subject string, body string) error {
	args := m.Called(to, subject, body)
	return args.Error(0)
}
