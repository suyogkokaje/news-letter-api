package repository

import (
	"errors"
	"gorm.io/gorm"

	news_letter_model "go_newsletter_api/internal/news_letter/model"
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

func (nr *NewsletterRepository) GetSubscribers(adminID uint) ([]string, error) {
	var emails []string

	var adminNewsletter news_letter_model.Newsletter
	if err := nr.DB.
		Where("admin_id = ?", adminID).
		First(&adminNewsletter).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("Admin does not have a newsletter")
		}
		return nil, err
	}

	if err := nr.DB.
		Table("newsletter_subscribers").
		Select("users.email").
		Joins("JOIN users ON users.id = newsletter_subscribers.user_id").
		Where("newsletter_subscribers.newsletter_id = ?", adminNewsletter.ID).
		Pluck("email", &emails).
		Error; err != nil {
		return nil, err
	}

	return emails, nil
}

func (nr *NewsletterRepository) FetchAllNewsletters() ([]news_letter_model.Newsletter, error) {
	var newsletters []news_letter_model.Newsletter
	err := nr.DB.Find(&newsletters).Error
	if err != nil {
		return nil, err
	}
	return newsletters, nil
}

func (nr *NewsletterRepository) DeleteNewsletter(newsletterID uint) error {
    var newsletter news_letter_model.Newsletter
    if err := nr.DB.Where("id = ?", newsletterID).First(&newsletter).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("Newsletter not found")
        }
        return err
    }

    if err := nr.DB.Delete(&newsletter).Error; err != nil {
        return err
    }

    return nil
}

func (nr *NewsletterRepository) UpdateNewsletter(newsletter *news_letter_model.Newsletter) error {
	var existingNewsletter news_letter_model.Newsletter
	if err := nr.DB.
		Where("id = ?", newsletter.ID).
		First(&existingNewsletter).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("Newsletter not found")
		}
		return err
	}

	if err := nr.DB.Save(newsletter).Error; err != nil {
		return err
	}

	return nil
}

