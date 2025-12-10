package models

type QuotationItem struct {
    ID          uint    `gorm:"primaryKey"`
    ItemName    string  `gorm:"not null"`
    Quantity    int     `gorm:"not null"`
    UnitPrice   float64 `gorm:"not null"`
    Total       float64 `gorm:"not null"`

    QuotationID uint
}
