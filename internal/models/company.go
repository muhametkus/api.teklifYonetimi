package models

import "time"

type Company struct {
    ID           uint            `gorm:"primaryKey"`
    Name         string          `gorm:"not null"`
    Logo         *string
    Subscription SubscriptionType `gorm:"type:varchar(20);default:'BASIC'"`
    UsersCount   int             `gorm:"default:0"`

    CreatedAt    time.Time
    UpdatedAt    time.Time

    Users       []User
    Quotations  []Quotation
}
