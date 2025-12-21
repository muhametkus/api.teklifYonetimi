package models

import "time"

type Quotation struct {
    ID          uint             `gorm:"primaryKey"`
    Title       string           `gorm:"not null"`
    Customer    string           `gorm:"not null"`
    Description *string
    Status      QuotationStatus  `gorm:"type:varchar(20);default:'PENDING'"`
    Total       float64          `gorm:"default:0"`

    CompanyID   uint
    Company     Company

    CreatedBy   uint

    Items       []QuotationItem

    CreatedAt   time.Time
    UpdatedAt   time.Time
}
