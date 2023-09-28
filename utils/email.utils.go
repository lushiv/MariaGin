package common_utils

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendGridMailHelper defines a helper for sending emails using SendGrid.
type SendGridMailHelper struct{}

// SendMail sends an email using SendGrid.
func (s *SendGridMailHelper) SendMail(message *MailMessage) (*SendGridMailResponse, error) {
	// Get the SendGrid API key
	sendgridAPIKey := os.Getenv("SENDGRID_API_KEY")

	// Return an error if API key is not set
	if sendgridAPIKey == "" {
		return nil, fmt.Errorf("Sendgrid API key not set")
	}

	// Get the SendGrid sender email address
	sendgridSenderEmail := os.Getenv("SENDGRID_SENDER_EMAIL")

	// Return an error if sender email is not set
	if sendgridSenderEmail == "" {
		return nil, fmt.Errorf("Sendgrid sender email address not set")
	}

	// Create a SendGrid client with the API key
	client := sendgrid.NewSendClient(sendgridAPIKey)

	// Prepare the email message
	messageBody := message.Body
	from := mail.NewEmail("Sender", sendgridSenderEmail)
	to := mail.NewEmail("Recipient", message.Email)
	subject := message.Title
	content := mail.NewContent("text/html", messageBody)
	email := mail.NewV3MailInit(from, subject, to, content)

	// Send the email
	response, err := client.Send(email)
	if err != nil {
		return nil, err
	}
	fmt.Println("response", response)

	return &SendGridMailResponse{
		Success: true,
		Content: "response",
	}, nil
}

// MailMessage represents an email message.
type MailMessage struct {
	Email string
	Title string
	Body  string
}

// SendGridMailResponse represents a SendGrid email response.
type SendGridMailResponse struct {
	Success bool
	Content string // Add this field if you want to store the content
}
