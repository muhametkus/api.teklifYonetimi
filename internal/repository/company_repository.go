package repository

import (
    "api.teklifYonetimi/internal/database"
    "api.teklifYonetimi/internal/models"
)

// CompanyRepository
// Bu struct DB işlemleri için kullanılır
type CompanyRepository struct {
}

// NewCompanyRepository
// Repository instance üretir
func NewCompanyRepository() *CompanyRepository {
    return &CompanyRepository{}
}

// Create
// Yeni company kaydı oluşturur
func (r *CompanyRepository) Create(company *models.Company) error {
    return database.DB.Create(company).Error
}

// FindAll
// Tüm company kayıtlarını getirir
func (r *CompanyRepository) FindAll() ([]models.Company, error) {
    var companies []models.Company
    err := database.DB.Find(&companies).Error
    return companies, err
}

// FindByID
// ID ile company getirir
func (r *CompanyRepository) FindByID(id uint) (*models.Company, error) {
    var company models.Company

    err := database.DB.First(&company, id).Error
    if err != nil {
        return nil, err
    }

    return &company, nil
}

// Update
// Company kaydını günceller
func (r *CompanyRepository) Update(company *models.Company) error {
    return database.DB.Save(company).Error
}

