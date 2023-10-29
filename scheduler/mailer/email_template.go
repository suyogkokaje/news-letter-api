package scheduler

import (
	"fmt"
	edition_model "go_newsletter_api/internal/edition/model"
	"time"
)

const (
	emailSubjectTemplate = "New Edition Published"
	emailBodyTemplate    = `Dear Subscriber,

A new edition has been published for the newsletter. Here are the details

Edition Title: %s
Publication Date: %s

Thank you for subscribing to our newsletter!`
)

func getEmailSubject() string {
	return emailSubjectTemplate
}

func getEmailBody(edition edition_model.Edition) string {
	return fmt.Sprintf(emailBodyTemplate, edition.Name, edition.PublishAt.Format(time.RFC3339))
}
