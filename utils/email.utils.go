package common_utils

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// EmailData represents the data required for sending an email.
type EmailData struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func SendEmailUsingSendGrid(data *EmailData) error {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	fmt.Println("SENDGRID_API_KEY:", apiKey)
	sgClient := sendgrid.NewSendClient(apiKey)

	// Log recipient email, subject, and message
	fmt.Printf("Recipient: %s\nSubject: %s\nMessage: %s\n", data.To, data.Subject, data.Message)

	// Create an email message
	from := mail.NewEmail("Sender Name", os.Getenv("SENDGRID_SENDER_EMAIL"))
	to := mail.NewEmail("Recipient Name", data.To)
	subject := data.Subject
	plainTextContent := data.Message
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, "")

	// Send the email
	response, err := sgClient.Send(message)
	if err != nil {
		fmt.Println("SendGrid Error:", err)
	} else {
		fmt.Println("SendGrid Response:", response.StatusCode)
		return err
	}
	return err
}
