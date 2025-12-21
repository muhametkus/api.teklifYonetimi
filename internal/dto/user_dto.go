package dto

// CreateUserRequest
// POST /users
type CreateUserRequest struct {
    Name      string `json:"name" binding:"required"`
    Email     string `json:"email" binding:"required,email"`
    Password  string `json:"password" binding:"required,min=6"`
    Role      string `json:"role"`
    CompanyID *uint  `json:"company_id"`
}

// UpdateUserRequest
// PUT /users/:id
type UpdateUserRequest struct {
    Name     string `json:"name"`
    Email    string `json:"email" binding:"omitempty,email"`
    Password string `json:"password" binding:"omitempty,min=6"`
    Role     string `json:"role"`
}
