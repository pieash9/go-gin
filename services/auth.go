package services

import (
	"errors"

	"github.com/pieash9/go-gin/internal/model"
	"github.com/pieash9/go-gin/internal/utils"
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

func (a *AuthService) CheckUserExistOrNot(email string) bool {
	var user model.User

	if err := a.db.Where("email=?", email).Find(&user).Error; err != nil {
		return false
	}

	if user.Email != "" {
		return true
	}

	return false
}

func (a *AuthService) Login(email *string, password *string) (*model.User, error) {
	if email == nil {
		return nil, errors.New("email is required")
	}
	if password == nil {
		return nil, errors.New("password is required")
	}

	var user model.User

	if err := a.db.Where("email=?", email).Find(&user).Error; err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(*password, user.Password) {
		return nil, errors.New("invalid password")
	}

	if user.Email == "" {
		return nil, errors.New("user not found with this email")
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

	if a.CheckUserExistOrNot(*email) {
		return nil, errors.New("user already exist with this email")
	}

	hashedPassword, err := utils.HashPassword(*password)
	if err != nil {
		return nil, err
	}

	var user model.User

	user.Email = *email
	user.Password = hashedPassword

	if err := a.db.Create(&user).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
