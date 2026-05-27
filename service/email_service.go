package service

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailService interface {
	SendEmail(to string, subject string, content string) error
}

type emailService struct {
	apiKey string
	sender string
}

func NewEmailService() EmailService {
	return &emailService{
		apiKey: os.Getenv("SENDGRID_API_KEY"),
		sender: os.Getenv("EMAIL_SENDER"),
	}
}

func (e *emailService) SendEmail(to string, subject string, content string) error {

	if e.apiKey == "" {
		return fmt.Errorf("sendgrid api key is empty")
	}

	if e.sender == "" {
		return fmt.Errorf("email sender is empty")
	}

	from := mail.NewEmail("Luxury Hotel Rental", e.sender)
	toEmail := mail.NewEmail("User", to)

	htmlContent := fmt.Sprintf(`
		<h2>%s</h2>
		<p>%s</p>
	`, subject, content)

	message := mail.NewSingleEmail(from, subject, toEmail, content, htmlContent)

	client := sendgrid.NewSendClient(e.apiKey)
	response, err := client.Send(message)

	if err != nil {
		return fmt.Errorf("sendgrid send error: %v", err)
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf(
			"sendgrid failed (status=%d, body=%s)",
			response.StatusCode,
			response.Body,
		)
	}

	return nil
}
