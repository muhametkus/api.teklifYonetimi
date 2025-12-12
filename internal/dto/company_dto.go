package dto

// CreateCompanyRequest
// POST /companies
type CreateCompanyRequest struct {
    Name string  `json:"name" binding:"required"`
    Logo *string `json:"logo"`
}

// UpdateCompanyRequest
// PUT /companies/:id
type UpdateCompanyRequest struct {
    Name         string  `json:"name"`
    Logo         *string `json:"logo"`
    Subscription string  `json:"subscription"`
}
