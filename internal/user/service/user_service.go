package service

import (
    news_letter_model "go_newsletter_api/internal/news_letter/model"
	user_model "go_newsletter_api/internal/user/model"
    "go_newsletter_api/internal/user/repository"
    "golang.org/x/crypto/bcrypt"
)

type UserService struct {
    UserRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{
        UserRepository: repo,
    }
}

func (us *UserService) CreateUser(user *user_model.User) (*user_model.User, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    user.Password = string(hashedPassword)

    return us.UserRepository.CreateUser(user)
}

func (us *UserService) GetUserByEmail(email string) (*user_model.User, error) {
    return us.UserRepository.GetUserByEmail(email)
}

func (us *UserService) VerifyPassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (us *UserService) GetUserSubscriptions(userID uint) ([]news_letter_model.Newsletter, error) {
    return us.UserRepository.GetUserSubscriptions(userID)
}
