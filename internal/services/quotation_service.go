package services

import (
	"api.teklifYonetimi/internal/models"
	"api.teklifYonetimi/internal/repository"
	"errors"
)

type QuotationService struct {
	repo *repository.QuotationRepository
}

func NewQuotationService(repo *repository.QuotationRepository) *QuotationService {
	return &QuotationService{
		repo: repo,
	}
}

// CreateQuotation
func (s *QuotationService) CreateQuotation(
	companyID uint,
	title string,
	customer string,
	description string,
	itemRequests []models.QuotationItem,
) (*models.Quotation, error) {

	var total float64
	var items []models.QuotationItem

	// 1️⃣ Item'ları hazırla + satır toplamı hesapla
	for _, item := range itemRequests {
		itemTotal := float64(item.Quantity) * item.UnitPrice
		item.Total = itemTotal

		total += itemTotal
		items = append(items, item)
	}

	// 2️⃣ Quotation oluştur
	quotation := &models.Quotation{
		Title:       title,
		Customer:    customer,
		Description: description,
		Status:      models.QuotationStatus("PENDING"),
		Total:       total,
		CompanyID:   companyID,
	}

	// 3️⃣ Repository ile transaction
	if err := s.repo.CreateQuotationWithItems(quotation, items); err != nil {
		return nil, err
	}

	return quotation, nil
}

// GetCompanyQuotations
func (s *QuotationService) GetCompanyQuotations(companyID uint) ([]models.Quotation, error) {
	return s.repo.FindAllByCompany(companyID)
}


// UpdateQuotationStatus
func (s *QuotationService) UpdateQuotationStatus(
	companyID uint,
	quotationID uint,
	newStatus string,
) error {

	// 1️⃣ Teklif bu company'e mi ait?
	quotations, err := s.repo.FindAllByCompany(companyID)
	if err != nil {
		return err
	}

	var found bool
	for _, q := range quotations {
		if q.ID == quotationID {
			found = true
			break
		}
	}

	if !found {
		return errors.New("teklif bulunamadı")
	}

	// 2️⃣ Status'u güncelle
	return s.repo.UpdateStatus(
		quotationID,
		models.QuotationStatus(newStatus),
	)
}

// GetCompanyQuotationsPaginated
func (s *QuotationService) GetCompanyQuotationsPaginated(
	companyID uint,
	page int,
	limit int,
) ([]models.Quotation, int64, error) {

	// güvenlik
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.repo.FindAllByCompanyPaginated(companyID, page, limit)
}
