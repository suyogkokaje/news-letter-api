package repository

import (
    "github.com/jinzhu/gorm"
    "go_newsletter_api/internal/user/model"
)

type UserRepository struct {
    DB *gorm.DB
}

func (ur *UserRepository) CreateUser(user *model.User) (*model.User, error) {
    if err := ur.DB.Create(user).Error; err != nil {
        return nil, err
    }
    return user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*model.User, error) {
    var user model.User
    if err := ur.DB.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
