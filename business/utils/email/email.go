package email

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"gopkg.in/gomail.v2"
)

type emailClient struct{}

type emailClientInterface interface {
	SendEmail(email string, data *EmailData, dir string, htmlTemplate string)
}

var (
	EmailClient emailClientInterface
)

func init() {
	EmailClient = &emailClient{}
}

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

// 👇 Email template parser

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func (c *emailClient) SendEmail(email string, data *EmailData, dir string, htmlTemplate string) {

	// Sender data.
	from := os.Getenv("EMAIL_FROM")
	smtpPass := os.Getenv("SMTP_PASS")
	smtpUser := os.Getenv("SMTP_USER")
	to := email
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	var body bytes.Buffer

	template, err := ParseTemplateDir(dir)
	if err != nil {
		log.Print("Could not parse template", err)
	}

	template.ExecuteTemplate(&body, htmlTemplate, &data)
	// template.ExecuteTemplate(&body, "verificationCode.html", &data)

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	// m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		log.Print("Could not send email: ", err)
	}

}
