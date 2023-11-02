package repository

import (
    "errors"
    edition_model "go_newsletter_api/internal/edition/model"
    newsletter_model "go_newsletter_api/internal/news_letter/model"

    "gorm.io/gorm"
)

type EditionRepository struct {
    DB *gorm.DB
}

func NewEditionRepository(db *gorm.DB) *EditionRepository {
    return &EditionRepository{DB: db}
}

func (er *EditionRepository) CreateEdition(edition *edition_model.Edition) error {
    var newsletter newsletter_model.Newsletter
    if err := er.DB.Where("id = ?", edition.NewsletterID).First(&newsletter).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return errors.New("Newsletter does not exist")
        }
        return err
    }
    return er.DB.Create(edition).Error
}

func (er *EditionRepository) UpdateEdition(edition *edition_model.Edition) error {
    return er.DB.Save(edition).Error
}

func (er *EditionRepository) GetEditionByID(id uint) (*edition_model.Edition, error) {
    var edition edition_model.Edition
    err := er.DB.Where("id = ?", id).First(&edition).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("Edition not found")
        }
        return nil, err
    }
    return &edition, nil
}

func (er *EditionRepository) GetEditionsByNewsletterID(newsletterID uint, page, pageSize int) ([]edition_model.Edition, int, error) {
    var editions []edition_model.Edition
    var count int64

    query := er.DB.Where("newsletter_id = ?", newsletterID)
    query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&editions)
    query.Model(&editions).Count(&count)

    return editions, int(count), nil
}

func (er *EditionRepository) DeleteEdition(id uint) error {
    return er.DB.Delete(&edition_model.Edition{}, id).Error
}
