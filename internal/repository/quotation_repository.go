package repository

import (
	"api.teklifYonetimi/internal/database"
	"api.teklifYonetimi/internal/models"

	"gorm.io/gorm"
)

type QuotationRepository struct {
	db *gorm.DB
}

func NewQuotationRepository() *QuotationRepository {
	return &QuotationRepository{
		db: database.DB,
	}
}

// CreateQuotationWithItems
// Transaction içinde quotation + items kaydeder
func (r *QuotationRepository) CreateQuotationWithItems(
	quotation *models.Quotation,
	items []models.QuotationItem,
) error {

	return r.db.Transaction(func(tx *gorm.DB) error {

		// 1️⃣ Quotation oluştur
		if err := tx.Create(quotation).Error; err != nil {
			return err
		}

		// 2️⃣ Item'ları quotation'a bağla
		for i := range items {
			items[i].QuotationID = quotation.ID
		}

		// 3️⃣ Item'ları kaydet
		if err := tx.Create(&items).Error; err != nil {
			return err
		}

		return nil
	})
}

// FindAllByCompany
func (r *QuotationRepository) FindAllByCompany(companyID uint) ([]models.Quotation, error) {
	var quotations []models.Quotation

	err := r.db.
		Where("company_id = ?", companyID).
		Preload("Items").
		Find(&quotations).Error

	return quotations, err
}

// UpdateStatus
func (r *QuotationRepository) UpdateStatus(
	quotationID uint,
	status models.QuotationStatus,
) error {

	return r.db.Model(&models.Quotation{}).
		Where("id = ?", quotationID).
		Update("status", status).Error
}

// FindAllByCompanyPaginated
func (r *QuotationRepository) FindAllByCompanyPaginated(
	companyID uint,
	page int,
	limit int,
) ([]models.Quotation, int64, error) {

	var quotations []models.Quotation
	var total int64

	offset := (page - 1) * limit

	// total count
	if err := r.db.
		Model(&models.Quotation{}).
		Where("company_id = ?", companyID).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// data
	err := r.db.
		Where("company_id = ?", companyID).
		Preload("Items").
		Limit(limit).
		Offset(offset).
		Order("created_at desc").
		Find(&quotations).Error

	return quotations, total, err
}
