package scheduler

import (
    "go_newsletter_api/config"
    edition_model "go_newsletter_api/internal/edition/model"
    "log"
    "os"
    "net/smtp"
    "time"
    "github.com/robfig/cron"
    edition_service "go_newsletter_api/internal/edition/service"
    news_letter_service "go_newsletter_api/internal/news_letter/service"
)

var failedEmails []string

func StartEditionPublishScheduler(editionService edition_service.EditionService, newsletterService news_letter_service.NewsletterService) {
    c := cron.New()
    _ = c.AddFunc("45 8 12 * * *", func() {
        log.Printf("Running The CRON JOB!!")
        newsletters, err := newsletterService.FetchAllNewsletters()
        if err != nil {
            log.Printf("Error fetching newsletters: %v", err)
            return
        }

        for _, newsletter := range newsletters {
            editions, err := editionService.GetEditionsByNewsletterID(newsletter.ID)
            if err != nil {
                log.Printf("Error fetching editions for newsletter %d: %v", newsletter.ID, err)
                continue
            }

            for _, edition := range editions {
                if edition.IsCompleted && !edition.IsPublished && edition.PublishAt.Before(time.Now()) {
                    edition.IsPublished = true

                    err := editionService.UpdateEdition(&edition)
                    if err != nil {
                        log.Printf("Error updating edition %d: %v", edition.ID, err)
                        continue
                    }

                    subscribers, err := newsletterService.GetSubscribers(newsletter.AdminID)
                    if err != nil {
                        log.Printf("Error fetching subscribers for newsletter %d: %v", newsletter.ID, err)
                        continue
                    }

                    for _, subscriber := range subscribers {
                        sendEditionDetailsToSubscriber(subscriber, edition)
                    }
                }
            }
        }
		log.Printf("The End!")
    })
    c.Start()
}

func sendEditionDetailsToSubscriber(subscriberEmail string, edition edition_model.Edition) {
    if err := config.LoadConfig(); err != nil {
        log.Fatal("Failed to load configuration")
    }

    senderEmail := os.Getenv("SENDER_EMAIL")
	secretKey := os.Getenv("SECRET_KEY")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	// log.Printf(senderEmail)

    emailSubject := "New Edition Published"
    emailBody := "Dear Subscriber,\n\n" +
        "A new edition has been published for the newsletter. Here are the details:\n\n" +
        "Edition Title: " + edition.Name + "\n" +
        "Publication Date: " + edition.PublishAt.Format(time.RFC3339) + "\n" +
        "\n\nThank you for subscribing to our newsletter!"

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

func PrintFailedEmails() {
    log.Println("Failed Emails:")
    for _, email := range failedEmails {
        log.Println(email)
    }
}
