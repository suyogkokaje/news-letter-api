package scheduler

import (
	edition_service "go_newsletter_api/internal/edition/service"
	news_letter_service "go_newsletter_api/internal/news_letter/service"
	"log"
	"time"

	"github.com/robfig/cron"
)

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
