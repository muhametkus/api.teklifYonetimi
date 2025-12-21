package repository

import (
    "api.teklifYonetimi/internal/database"
    "api.teklifYonetimi/internal/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
    return &UserRepository{}
}

// Create
func (r *UserRepository) Create(user *models.User) error {
    return database.DB.Create(user).Error
}

// FindAll
func (r *UserRepository) FindAll() ([]models.User, error) {
    var users []models.User
    err := database.DB.Find(&users).Error
    return users, err
}

// FindByCompanyID
func (r *UserRepository) FindByCompanyID(companyID uint) ([]models.User, error) {
    var users []models.User
    err := database.DB.Where("company_id = ?", companyID).Find(&users).Error
    return users, err
}

// FindByEmail
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
    var user models.User
    err := database.DB.Where("email = ?", email).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

// FindByID
func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := database.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update
func (r *UserRepository) Update(user *models.User) error {
    return database.DB.Save(user).Error
}

// Delete
func (r *UserRepository) Delete(id uint) error {
    return database.DB.Delete(&models.User{}, id).Error
}
