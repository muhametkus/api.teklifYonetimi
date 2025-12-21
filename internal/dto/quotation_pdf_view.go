package dto

import "api.teklifYonetimi/internal/models"

type QuotationPDFView struct {
	Title       string
	Customer    string
	Status      string
	Total       float64
	Items       []models.QuotationItem

	CompanyName string
	CompanyLogo string
}
