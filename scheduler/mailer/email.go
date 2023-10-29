package scheduler

import (
	"go_newsletter_api/config"
	edition_model "go_newsletter_api/internal/edition/model"
	"log"
	"net/smtp"
	"os"
)

var failedEmails []string

func sendEditionDetailsToSubscriber(subscriberEmail string, edition edition_model.Edition) {
	if err := config.LoadConfig(); err != nil {
		log.Fatal("Failed to load configuration")
	}

	senderEmail := os.Getenv("SENDER_EMAIL")
	secretKey := os.Getenv("SECRET_KEY")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	emailSubject := getEmailSubject()
	emailBody := getEmailBody(edition)

	auth := smtp.PlainAuth("", senderEmail, secretKey, smtpHost)

	msg := []byte("To: " + subscriberEmail + "\r\n" +
		"Subject: " + emailSubject + "\r\n" +
		"\r\n" +
		emailBody)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{subscriberEmail}, msg)
	if err != nil {
		log.Printf(subscriberEmail)
		log.Printf("Error sending email to subscriber %s: %v", subscriberEmail, err)
		failedEmails = append(failedEmails, subscriberEmail)
	} else {
		log.Printf("Email sent to subscriber %s", subscriberEmail)
	}
}
