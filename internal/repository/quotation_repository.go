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

// FindPaginatedByCompany
func (r *QuotationRepository) FindPaginatedByCompany(
	companyID uint,
	limit int,
	offset int,
) ([]models.Quotation, int64, error) {

	var quotations []models.Quotation
	var total int64

	// total count
	err := r.db.Model(&models.Quotation{}).
		Where("company_id = ?", companyID).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// paginated data
	err = r.db.
		Where("company_id = ?", companyID).
		Preload("Items").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&quotations).Error

	if err != nil {
		return nil, 0, err
	}

	return quotations, total, nil
}

// FindPaginatedByCompanyAndUser
func (r *QuotationRepository) FindPaginatedByCompanyAndUser(
	companyID uint,
	userID uint,
	limit int,
	offset int,
) ([]models.Quotation, int64, error) {

	var quotations []models.Quotation
	var total int64

	err := r.db.Model(&models.Quotation{}).
		Where("company_id = ? AND created_by = ?", companyID, userID).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.
		Where("company_id = ? AND created_by = ?", companyID, userID).
		Preload("Items").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&quotations).Error

	if err != nil {
		return nil, 0, err
	}

	return quotations, total, nil
}


// FindPaginatedFilteredByCompany
func (r *QuotationRepository) FindPaginatedFilteredByCompany(
	companyID uint,
	status string,
	customer string,
	limit int,
	offset int,
) ([]models.Quotation, int64, error) {

	var quotations []models.Quotation
	var total int64

	query := r.db.Model(&models.Quotation{}).
		Where("company_id = ?", companyID)

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if customer != "" {
		query = query.Where("customer ILIKE ?", "%"+customer+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.
		Preload("Items").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&quotations).Error

	if err != nil {
		return nil, 0, err
	}

	return quotations, total, nil
}

// FindPaginatedFilteredByCompanyAndUser
func (r *QuotationRepository) FindPaginatedFilteredByCompanyAndUser(
	companyID uint,
	userID uint,
	status string,
	customer string,
	limit int,
	offset int,
) ([]models.Quotation, int64, error) {

	var quotations []models.Quotation
	var total int64

	query := r.db.Model(&models.Quotation{}).
		Where("company_id = ? AND created_by = ?", companyID, userID)

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if customer != "" {
		query = query.Where("customer ILIKE ?", "%"+customer+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.
		Preload("Items").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&quotations).Error

	if err != nil {
		return nil, 0, err
	}

	return quotations, total, nil
}

// Create
func (r *QuotationRepository) Create(quotation *models.Quotation) error {
	return r.db.Create(quotation).Error
}

// FindByCompanyAndUser
func (r *QuotationRepository) FindByCompanyAndUser(companyID, userID uint) ([]models.Quotation, error) {
	var quotations []models.Quotation
	err := r.db.
		Where("company_id = ? AND created_by = ?", companyID, userID).
		Preload("Items").
		Order("created_at desc").
		Find(&quotations).Error
	return quotations, err
}

// FindByIDWithItems
func (r *QuotationRepository) FindByIDWithItems(id, companyID uint) (*models.Quotation, error) {
	var quotation models.Quotation
	err := r.db.
		Where("id = ? AND company_id = ?", id, companyID).
		Preload("Items").
		First(&quotation).Error
	if err != nil {
		return nil, err
	}
	return &quotation, nil
}

// Update
func (r *QuotationRepository) Update(quotation *models.Quotation) error {
    return r.db.Save(quotation).Error
}

// Delete
func (r *QuotationRepository) Delete(id uint) error {
    return r.db.Delete(&models.Quotation{}, id).Error
}

// DeleteItemsByQuotationID
func (r *QuotationRepository) DeleteItemsByQuotationID(quotationID uint) error {
    return r.db.Where("quotation_id = ?", quotationID).Delete(&models.QuotationItem{}).Error
}
