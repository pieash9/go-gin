package services

import (
	"errors"

	"github.com/pieash9/go-gin/internal/model"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func InitAuthService(db *gorm.DB) *AuthService {
	db.AutoMigrate(model.User{})
	return &AuthService{
		db: db,
	}
}

func (a *AuthService) Login(email *string, password *string) (*model.User, error) {
	if email == nil {
		return nil, errors.New("email is required")
	}
	if password == nil {
		return nil, errors.New("password is required")
	}

	var user model.User

	if err := a.db.Where("email=?", email).Where("password=?", password).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *AuthService) Register(email *string, password *string) (*model.User, error) {
	if email == nil {
		return nil, errors.New("email is required")
	}
	if password == nil {
		return nil, errors.New("password is required")
	}

	var user model.User

	user.Email = *email
	user.Password = *password

	if err := a.db.Create(&user).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
