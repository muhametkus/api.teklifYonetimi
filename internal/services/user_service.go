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

// GetUserByID
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
    return s.repo.FindByID(id)
}

// UpdateUser
func (s *UserService) UpdateUser(id uint, name, email, password, role string) (*models.User, error) {
    user, err := s.repo.FindByID(id)
    if err != nil {
        return nil, err
    }

    if name != "" {
        user.Name = name
    }
    if email != "" {
        // Check if email is taken by another user
        existingUser, err := s.repo.FindByEmail(email)
        if err == nil && existingUser.ID != id {
            return nil, errors.New("email zaten kullanılıyor")
        }
        user.Email = email
    }
    if password != "" {
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err != nil {
            return nil, err
        }
        user.Password = string(hashedPassword)
    }
    if role != "" {
        user.Role = models.UserRole(role)
    }

    if err := s.repo.Update(user); err != nil {
        return nil, err
    }

    return user, nil
}

// DeleteUser
func (s *UserService) DeleteUser(id uint) error {
    return s.repo.Delete(id)
}
