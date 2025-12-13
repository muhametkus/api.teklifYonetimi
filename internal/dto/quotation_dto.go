package dto

// QuotationItemRequest
type QuotationItemRequest struct {
	ItemName  string  `json:"item_name" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,min=1"`
	UnitPrice float64 `json:"unit_price" binding:"required,min=0"`
}

// CreateQuotationRequest
type CreateQuotationRequest struct {
	Title       string                 `json:"title" binding:"required"`
	Customer    string                 `json:"customer" binding:"required"`
	Description string                 `json:"description"`
	Items       []QuotationItemRequest `json:"items" binding:"required,min=1"`
}

// UpdateQuotationStatusRequest
type UpdateQuotationStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=APPROVED REJECTED"`
}
