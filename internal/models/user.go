package models

import "time"

type User struct {
    ID        uint     `gorm:"primaryKey"`
    Name      string
    Email     string    `gorm:"uniqueIndex;not null"`
    Password  string    `gorm:"not null"`
    Role      UserRole  `gorm:"type:varchar(20);default:'USER'"`
    Avatar    *string

    CompanyID *uint
    Company   *Company

    CreatedAt time.Time
    UpdatedAt time.Time
}
