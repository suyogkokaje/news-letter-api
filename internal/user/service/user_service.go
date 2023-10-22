package service

import (
    "go_newsletter_api/internal/user/model"
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

func (us *UserService) CreateUser(user *model.User) (*model.User, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    user.Password = string(hashedPassword)

    return us.UserRepository.CreateUser(user)
}

func (us *UserService) GetUserByEmail(email string) (*model.User, error) {
    return us.UserRepository.GetUserByEmail(email)
}

func (us *UserService) VerifyPassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
