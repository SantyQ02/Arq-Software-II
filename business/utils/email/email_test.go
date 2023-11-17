package email

import (
	"mvc-go/utils/initializers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	initializers.LoadTestEnv("../initializers/test.env")
}

func initTestClient() {
	EmailClient = &emailClient{}
}

func TestParseTemplateDir(t *testing.T) {
	_, err := ParseTemplateDir("../../templates")

	assert.Nil(t, err)
}

func TestEmail(t *testing.T) {
	initTestClient()
	emaildata := EmailData{
		FirstName: "Giovanni",
		URL:       "test.com/test/url",
		Subject:   "Email test",
	}

	EmailClient.SendEmail("testing@itvirtuous.com", &emaildata, "../../templates/resetPassword", "resetPasswordLink.html")
	EmailClient.SendEmail("testing@itvirtuous.com", &emaildata, "../../templates/verificateAccount", "verificateAccountLink.html")

}
