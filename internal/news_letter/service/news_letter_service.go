package service

import (
	"go_newsletter_api/internal/news_letter/model"
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

func (ns *NewsletterService) CreateNewsletter(newsletter *model.Newsletter, adminID uint) (*model.Newsletter, error) {
	return ns.NewsletterRepository.CreateNewsletter(newsletter, adminID)
}

func (ns *NewsletterService) SubscribeUser(newsletterID, userID uint) error {
	return ns.NewsletterRepository.SubscribeUser(newsletterID, userID)
}

func (ns *NewsletterService) UnsubscribeUser(newsletterID, userID uint) error {
	return ns.NewsletterRepository.UnsubscribeUser(newsletterID, userID)
}
