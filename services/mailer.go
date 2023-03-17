package services

import (
	"bytes"
	"html/template"

	"github.com/abgeo/domains/config"
	"github.com/pkg/errors"
	gomail "gopkg.in/mail.v2"
)

type MailerService struct {
	config *config.Config
}

func NewMailerService(conf *config.Config) *MailerService {
	return &MailerService{config: conf}
}

func (service *MailerService) Send(from string, templateName string, data any) error {
	tmpl, err := template.ParseFiles("./templates/" + templateName + ".html")
	if err != nil {
		return errors.Wrap(err, "unable to parse email template")
	}

	var tpl bytes.Buffer
	if err = tmpl.Execute(&tpl, data); err != nil {
		return errors.Wrap(err, "unable to execute email template")
	}

	message := gomail.NewMessage()

	message.SetHeader("From", service.config.Email.From)
	message.SetHeader("Reply-to", from)
	message.SetHeader("To", service.config.Email.To)
	message.SetHeader("Subject", service.config.Email.Subject)
	message.SetBody("text/html", tpl.String())

	dialer := gomail.NewDialer(
		service.config.SMTP.Host,
		service.config.SMTP.Port,
		service.config.SMTP.Username,
		service.config.SMTP.Password,
	)

	if err = dialer.DialAndSend(message); err != nil {
		return errors.Wrap(err, "unable to dial and send")
	}

	return nil
}
