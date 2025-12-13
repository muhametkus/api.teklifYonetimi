package services

import (
	"errors"

	"api.teklifYonetimi/internal/models"
	"api.teklifYonetimi/internal/repository"
	"api.teklifYonetimi/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

// Login
func (s *AuthService) Login(email string, password string) (string, *models.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", nil, errors.New("kullanıcı bulunamadı")
	}

	// Password check
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)
	if err != nil {
		return "", nil, errors.New("şifre hatalı")
	}

	// JWT üret
	token, err := utils.GenerateToken(user.ID, string(user.Role))
	if err != nil {
		return "", nil, err
	}

	// response'ta password dönmesin
	user.Password = ""

	return token, user, nil
}
