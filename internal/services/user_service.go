package services

import (
    "errors"

    "api.teklifYonetimi/internal/models"
    "api.teklifYonetimi/internal/repository"

    "golang.org/x/crypto/bcrypt"
)

// UserService
// İş mantığı burada
type UserService struct {
    repo *repository.UserRepository
}

// NewUserService
func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{
        repo: repo,
    }
}

// CreateUser
// Yeni user oluşturur (password hash + role default)
func (s *UserService) CreateUser(
    name string,
    email string,
    password string,
    role string,
    companyID *uint,
) (*models.User, error) {

    // 1️⃣ Email daha önce kullanılmış mı?
    _, err := s.repo.FindByEmail(email)
    if err == nil {
        return nil, errors.New("email zaten kullanılıyor")
    }

    // 2️⃣ Password hashle
    hashedPassword, err := bcrypt.GenerateFromPassword(
        []byte(password),
        bcrypt.DefaultCost,
    )
    if err != nil {
        return nil, err
    }

    // 3️⃣ Role default
    userRole := models.RoleUser
    if role != "" {
        userRole = models.UserRole(role)
    }

    // 4️⃣ User modelini oluştur
    user := &models.User{
        Name:      name,
        Email:     email,
        Password:  string(hashedPassword),
        Role:      userRole,
        CompanyID: companyID,
    }

    // 5️⃣ DB’ye kaydet
    if err := s.repo.Create(user); err != nil {
        return nil, err
    }

    return user, nil
}

// GetAllUsers
func (s *UserService) GetAllUsers() ([]models.User, error) {
    return s.repo.FindAll()
}
