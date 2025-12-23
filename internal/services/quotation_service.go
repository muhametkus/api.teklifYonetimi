package services

import (
	"errors"

	"api.teklifYonetimi/internal/models"
	"api.teklifYonetimi/internal/repository"
)

type QuotationService struct {
	repo *repository.QuotationRepository
}

// Constructor
func NewQuotationService(repo *repository.QuotationRepository) *QuotationService {
	return &QuotationService{
		repo: repo,
	}
}

//
// CREATE QUOTATION
//

func (s *QuotationService) CreateQuotation(
	companyID uint,
	userID uint,
	title string,
	customer string,
	description string,
	items []models.QuotationItem,
) (*models.Quotation, error) {

	if len(items) == 0 {
		return nil, errors.New("en az bir kalem eklenmelidir")
	}

	var total float64
	for i := range items {
		items[i].Total = float64(items[i].Quantity) * items[i].UnitPrice
		total += items[i].Total
	}

	quotation := &models.Quotation{
		Title:       title,
		Customer:    customer,
		Description: &description,
		Status:      models.QuotationStatus("PENDING"),
		Total:       total,
		CompanyID:   companyID,
		CreatedBy:   userID,
		Items:       items,
	}

	if err := s.repo.Create(quotation); err != nil {
		return nil, err
	}

	return quotation, nil
}

//
// LIST QUOTATIONS (ROLE BASED)
//

func (s *QuotationService) GetQuotationsByRole(
	companyID uint,
	userID uint,
	role string,
) ([]models.Quotation, error) {

	// ADMIN/SUPER_ADMIN → company'deki tüm teklifler
	if role == "ADMIN" || role == "SUPER_ADMIN" {
		return s.repo.FindAllByCompany(companyID)
	}

	// USER → sadece kendi oluşturdukları
	return s.repo.FindByCompanyAndUser(companyID, userID)
}

//
// GET SINGLE QUOTATION (PDF / DETAIL)
//

func (s *QuotationService) GetQuotationForPDF(
	companyID uint,
	quotationID uint,
) (*models.Quotation, error) {

	return s.repo.FindByIDWithItems(
		quotationID,
		companyID,
	)
}

// GetFilteredPaginatedQuotations
func (s *QuotationService) GetFilteredPaginatedQuotations(
	companyID uint,
	userID uint,
	role string,
	status string,
	customer string,
	page int,
	limit int,
) ([]models.Quotation, int64, error) {

	if page < 1 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	if role == "ADMIN" || role == "SUPER_ADMIN" {
		return s.repo.FindPaginatedFilteredByCompany(
			companyID,
			status,
			customer,
			limit,
			offset,
		)
	}

	return s.repo.FindPaginatedFilteredByCompanyAndUser(
		companyID,
		userID,
		status,
		customer,
		limit,
		offset,
	)
}

// UpdateQuotationStatus
func (s *QuotationService) UpdateQuotationStatus(
	companyID uint,
	quotationID uint,
	status string,
) error {
	return s.repo.UpdateStatus(quotationID, models.QuotationStatus(status))
}

// GetQuotationByID
func (s *QuotationService) GetQuotationByID(id, companyID uint) (*models.Quotation, error) {
    return s.repo.FindByIDWithItems(id, companyID)
}

// UpdateQuotation
func (s *QuotationService) UpdateQuotation(
    id, companyID, userID uint,
    role string,
    title, customer, description string,
    items []models.QuotationItem,
) (*models.Quotation, error) {
    
    quotation, err := s.repo.FindByIDWithItems(id, companyID)
    if err != nil {
        return nil, err
    }

    // Authorization: Only Owner or Admin can update
    if role != "ADMIN" && role != "SUPER_ADMIN" && quotation.CreatedBy != userID {
        return nil, errors.New("bu teklifi güncelleme yetkiniz yok")
    }

    // Update fields
    if title != "" {
        quotation.Title = title
    }
    if customer != "" {
        quotation.Customer = customer
    }
    if description != "" {
        quotation.Description = &description
    }

    // Update Items (Re-calculate total)
    if len(items) > 0 {
        // Delete old items
        if err := s.repo.DeleteItemsByQuotationID(id); err != nil {
            return nil, err
        }
        
        var total float64
        for i := range items {
            items[i].QuotationID = id
            items[i].Total = float64(items[i].Quantity) * items[i].UnitPrice
            total += items[i].Total
        }
        quotation.Items = items
        quotation.Total = total
    }

    if err := s.repo.Update(quotation); err != nil {
        return nil, err
    }

    return quotation, nil
}

// DeleteQuotation
func (s *QuotationService) DeleteQuotation(id, companyID, userID uint, role string) error {
    quotation, err := s.repo.FindByIDWithItems(id, companyID)
    if err != nil {
        return err
    }

    // Authorization: Only Owner or Admin can delete
    if role != "ADMIN" && role != "SUPER_ADMIN" && quotation.CreatedBy != userID {
        return errors.New("bu teklifi silme yetkiniz yok")
    }

    // Delete items first (or rely on cascade if configured, but explicit is safer here)
    if err := s.repo.DeleteItemsByQuotationID(id); err != nil {
        return err
    }

    return s.repo.Delete(id)
}
