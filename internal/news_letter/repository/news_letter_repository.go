package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	news_letter_model "go_newsletter_api/internal/news_letter/model"
	user_model "go_newsletter_api/internal/user/model"
)

type NewsletterRepository struct {
	DB *gorm.DB
}

func (nr *NewsletterRepository) CreateNewsletter(newsletter *news_letter_model.Newsletter, adminID uint) (*news_letter_model.Newsletter, error) {
	var existingNewsletter news_letter_model.Newsletter
	if err := nr.DB.
		Where("admin_id = ?", adminID).
		First(&existingNewsletter).Error; err == nil {
		return nil, errors.New("Admin already has a newsletter")
	} else if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	newsletter.AdminID = adminID

	if err := nr.DB.Create(newsletter).Error; err != nil {
		return nil, err
	}

	return newsletter, nil
}

func (nr *NewsletterRepository) SubscribeUser(newsletterID, userID uint) error {
	var existingNewsletter news_letter_model.Newsletter
	if err := nr.DB.
		Where("id = ?", newsletterID).
		First(&existingNewsletter).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("Newsletter does not exist")
		}
		return err
	}

	var existingSubscription news_letter_model.NewsletterSubscriber
	if err := nr.DB.
		Where("newsletter_id = ? AND user_id = ?", newsletterID, userID).
		First(&existingSubscription).Error; err == nil {
		return errors.New("User is already subscribed to this newsletter")
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	subscription := news_letter_model.NewsletterSubscriber{
		NewsletterID: newsletterID,
		UserID:       userID,
	}

	if err := nr.DB.Create(&subscription).Error; err != nil {
		return err
	}

	return nil
}

func (nr *NewsletterRepository) UnsubscribeUser(newsletterID, userID uint) error {
	var existingNewsletter news_letter_model.Newsletter
	if err := nr.DB.
		Where("id = ?", newsletterID).
		First(&existingNewsletter).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("Newsletter does not exist")
		}
		return err
	}

	var existingSubscription news_letter_model.NewsletterSubscriber
	if err := nr.DB.
		Where("newsletter_id = ? AND user_id = ?", newsletterID, userID).
		First(&existingSubscription).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("User is not subscribed to this newsletter")
		}
		return err
	}

	if err := nr.DB.
		Where("newsletter_id = ? AND user_id = ?", newsletterID, userID).
		Delete(news_letter_model.NewsletterSubscriber{}).Error; err != nil {
		return err
	}

	return nil
}


func (nr *NewsletterRepository) GetSubscribers(newsletterID uint) ([]user_model.User, error) {
	var subscribers []user_model.User

	if err := nr.DB.
		Model(&news_letter_model.Newsletter{}).
		Where("id = ?", newsletterID).
		Association("Subscribers").
		Find(&subscribers).Error; err != nil {
		return nil, err
	}

	return subscribers, nil
}
