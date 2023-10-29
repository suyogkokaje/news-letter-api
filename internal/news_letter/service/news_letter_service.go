package service

import (
	news_letter_model "go_newsletter_api/internal/news_letter/model"

	"go_newsletter_api/internal/news_letter/repository"
)

type NewsletterService struct {
	NewsletterRepository repository.NewsletterRepository
}

func NewNewsletterService(repo repository.NewsletterRepository) *NewsletterService {
	return &NewsletterService{
		NewsletterRepository: repo,
	}
}

func (ns *NewsletterService) CreateNewsletter(newsletter *news_letter_model.Newsletter, adminID uint) (*news_letter_model.Newsletter, error) {
	return ns.NewsletterRepository.CreateNewsletter(newsletter, adminID)
}

func (ns *NewsletterService) SubscribeUser(newsletterID, userID uint) error {
	return ns.NewsletterRepository.SubscribeUser(newsletterID, userID)
}

func (ns *NewsletterService) UnsubscribeUser(newsletterID, userID uint) error {
	return ns.NewsletterRepository.UnsubscribeUser(newsletterID, userID)
}

func (ns *NewsletterService) GetSubscribers(adminID uint) ([]string, error) {
	return ns.NewsletterRepository.GetSubscribers(adminID)
}

func (ns *NewsletterService) FetchAllNewsletters() ([]news_letter_model.Newsletter, error) {
	return ns.NewsletterRepository.FetchAllNewsletters()
}

func (ns *NewsletterService) DeleteNewsletter(newsletterID uint) error {
    return ns.NewsletterRepository.DeleteNewsletter(newsletterID)
}