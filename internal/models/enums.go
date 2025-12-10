package models

type UserRole string
const (
    RoleSuperAdmin UserRole = "SUPER_ADMIN"
    RoleAdmin      UserRole = "ADMIN"
    RoleUser       UserRole = "USER"
)

type SubscriptionType string
const (
    SubBasic    SubscriptionType = "BASIC"
    SubPro      SubscriptionType = "PRO"
    SubBusiness SubscriptionType = "BUSINESS"
)

type QuotationStatus string
const (
    StatusPending  QuotationStatus = "PENDING"
    StatusApproved QuotationStatus = "APPROVED"
    StatusRejected QuotationStatus = "REJECTED"
)
