package email

import (
	"github.com/stretchr/testify/mock"
)

type EmailMockClient struct {
	mock.Mock
}

func (c *EmailMockClient) SendEmail(email string, data *EmailData, dir string, htmlTemplate string) {
	c.Called(email, data, dir)
	return
}
