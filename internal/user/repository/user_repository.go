package repository

import (
	news_letter_model "go_newsletter_api/internal/news_letter/model"
	user_model "go_newsletter_api/internal/user/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (ur *UserRepository) CreateUser(user *user_model.User) (*user_model.User, error) {
	if err := ur.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*user_model.User, error) {
	var user user_model.User
	if err := ur.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) GetUserSubscriptions(userID uint) ([]news_letter_model.Newsletter, error) {
	var newsletters []news_letter_model.Newsletter

	if err := ur.DB.Table("newsletter_subscribers").
		Select("newsletters.*").
		Joins("JOIN newsletters ON newsletters.id = newsletter_subscribers.newsletter_id").
		Where("newsletter_subscribers.user_id = ?", userID).
		Scan(&newsletters).
		Error; err != nil {
		return nil, err
	}

	return newsletters, nil
}
