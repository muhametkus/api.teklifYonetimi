package services

import (
    "api.teklifYonetimi/internal/models"
    "api.teklifYonetimi/internal/repository"
)

// CompanyService
// İş mantığı burada olur
type CompanyService struct {
    repo *repository.CompanyRepository
}

// NewCompanyService
// Service instance üretir
func NewCompanyService(repo *repository.CompanyRepository) *CompanyService {
    return &CompanyService{
        repo: repo,
    }
}

// CreateCompany
// Yeni company oluşturur
func (s *CompanyService) CreateCompany(name string, logo *string) (*models.Company, error) {
    company := &models.Company{
        Name: name,
        Logo: logo,
    }

    if err := s.repo.Create(company); err != nil {
        return nil, err
    }

    return company, nil
}

// GetAllCompanies
// Tüm company listesini getirir
func (s *CompanyService) GetAllCompanies() ([]models.Company, error) {
    return s.repo.FindAll()
}

// GetCompanyByID
// ID ile company getirir
func (s *CompanyService) GetCompanyByID(id uint) (*models.Company, error) {
    return s.repo.FindByID(id)
}


// UpdateCompany
// Company bilgilerini günceller
func (s *CompanyService) UpdateCompany(
    id uint,
    name string,
    logo *string,
    subscription string,
) (*models.Company, error) {

    // 1️⃣ Önce company var mı kontrol et
    company, err := s.repo.FindByID(id)
    if err != nil {
        return nil, err
    }

    // 2️⃣ Alanları güncelle (boş gelmeyenler)
    if name != "" {
        company.Name = name
    }

    if logo != nil {
        company.Logo = logo
    }

    if subscription != "" {
        company.Subscription = models.SubscriptionType(subscription)
    }

    // 3️⃣ DB’ye kaydet
    if err := s.repo.Update(company); err != nil {
        return nil, err
    }

    return company, nil
}
