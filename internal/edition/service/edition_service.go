package service

import (
	"go_newsletter_api/internal/edition/model"
	"go_newsletter_api/internal/edition/repository"
)

type EditionService struct {
	EditionRepository repository.EditionRepository
}

func NewEditionService(repo repository.EditionRepository) *EditionService {
	return &EditionService{
		EditionRepository: repo,
	}
}

func (es *EditionService) CreateEdition(edition *model.Edition) error {
	return es.EditionRepository.CreateEdition(edition)
}

func (es *EditionService) UpdateEdition(edition *model.Edition) error {
	return es.EditionRepository.UpdateEdition(edition)
}

func (es *EditionService) GetEditionByID(id uint) (*model.Edition, error) {
	return es.EditionRepository.GetEditionByID(id)
}

func (es *EditionService) GetEditionsByNewsletterID(newsletterID uint) ([]model.Edition, error) {
	return es.EditionRepository.GetEditionsByNewsletterID(newsletterID)
}

func (es *EditionService) DeleteEdition(id uint) error {
	return es.EditionRepository.DeleteEdition(id)
}
