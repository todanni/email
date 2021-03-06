package email

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	log "github.com/sirupsen/logrus"
)

// TODO: Implement retrieve template by name using the Admin API
const (
	VerificationTemplateID  = "d-35759ae691be4ee28aced5ca730e50b5"
	PasswordResetTemplateID = "d-b64d65d60a674eed87454cd98eca1adc"
)

var (
	AdminSender = mail.NewEmail("ToDanni admin", "admin@todanni.com")
)

type emailService struct {
	client *sendgrid.Client
}

func NewEmailService(key string) Service {
	service := &emailService{
		client: sendgrid.NewSendClient(key),
	}
	return service
}

func (e emailService) SendPasswordResetEmail(code string, recipient Recipient) error {
	return e.sendEmail(PasswordResetTemplateID, code, recipient)
}

func (e emailService) SendVerificationEmail(code string, recipient Recipient) error {
	return e.sendEmail(VerificationTemplateID, code, recipient)
}

func (e emailService) sendEmail(templateID, link string, recipient Recipient) error {
	// Compose email
	message := mail.NewV3Mail()
	message.SetFrom(AdminSender)
	message.SetTemplateID(templateID)

	// Add recipient
	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail(recipient.FullName, recipient.Email),
	}
	p.AddTos(tos...)

	// Set verification code
	p.SetDynamicTemplateData("VerificationLink", link)
	message.AddPersonalizations(p)

	response, err := e.client.Send(message)
	if err != nil {
		return err
	}
	log.Info(response.StatusCode)
	return nil
}
