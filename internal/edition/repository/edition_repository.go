package repository

import (
	"github.com/jinzhu/gorm"
	edition_model "go_newsletter_api/internal/edition/model"
)

type EditionRepository struct {
	DB *gorm.DB
}

func NewEditionRepository(db *gorm.DB) *EditionRepository {
	return &EditionRepository{DB: db}
}

func (er *EditionRepository) CreateEdition(edition *edition_model.Edition) error {
	return er.DB.Create(edition).Error
}

func (er *EditionRepository) UpdateEdition(edition *edition_model.Edition) error {
	return er.DB.Save(edition).Error
}

func (er *EditionRepository) GetEditionByID(id uint) (*edition_model.Edition, error) {
	var edition edition_model.Edition
	err := er.DB.Where("id = ?", id).First(&edition).Error
	if err != nil {
		return nil, err
	}
	return &edition, nil
}

func (er *EditionRepository) GetEditionsByNewsletterID(newsletterID uint) ([]edition_model.Edition, error) {
	var editions []edition_model.Edition
	err := er.DB.Where("newsletter_id = ?", newsletterID).Find(&editions).Error
	if err != nil {
		return nil, err
	}
	return editions, nil
}

func (er *EditionRepository) DeleteEdition(id uint) error {
	return er.DB.Delete(&edition_model.Edition{}, id).Error
}
