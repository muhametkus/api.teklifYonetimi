package handlers

import (
    "net/http"
    "strconv"

    "api.teklifYonetimi/internal/dto"
    "api.teklifYonetimi/internal/repository"
    "api.teklifYonetimi/internal/services"

    "github.com/gin-gonic/gin"
)

// CompanyHandler
// HTTP layer
type CompanyHandler struct {
    service *services.CompanyService
}

// NewCompanyHandler
// Handler instance oluşturur
func NewCompanyHandler() *CompanyHandler {
    repo := repository.NewCompanyRepository()
    service := services.NewCompanyService(repo)

    return &CompanyHandler{
        service: service,
    }
}

// CreateCompany
// POST /companies
func (h *CompanyHandler) CreateCompany(c *gin.Context) {
    var req dto.CreateCompanyRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   err.Error(),
        })
        return
    }

    company, err := h.service.CreateCompany(req.Name, req.Logo)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "error":   "Company oluşturulamadı",
        })
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "success": true,
        "data":    company,
    })
}

// GetCompanies
// GET /companies
func (h *CompanyHandler) GetCompanies(c *gin.Context) {
    companies, err := h.service.GetAllCompanies()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "error":   "Company listesi alınamadı",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "data":    companies,
    })
}

// GetCompanyByID
// GET /companies/:id
func (h *CompanyHandler) GetCompanyByID(c *gin.Context) {
    idParam := c.Param("id")

    id, err := strconv.ParseUint(idParam, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "Geçersiz company id",
        })
        return
    }

    company, err := h.service.GetCompanyByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "success": false,
            "error":   "Company bulunamadı",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "data":    company,
    })
}

// UpdateCompany
// PUT /companies/:id
func (h *CompanyHandler) UpdateCompany(c *gin.Context) {
    idParam := c.Param("id")

    id, err := strconv.ParseUint(idParam, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "Geçersiz company id",
        })
        return
    }

    var req dto.UpdateCompanyRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   err.Error(),
        })
        return
    }

    company, err := h.service.UpdateCompany(
        uint(id),
        req.Name,
        req.Logo,
        req.Subscription,
    )

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "success": false,
            "error":   "Company bulunamadı veya güncellenemedi",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "data":    company,
    })
}


// DeleteCompany
// DELETE /companies/:id
func (h *CompanyHandler) DeleteCompany(c *gin.Context) {
    idParam := c.Param("id")

    id, err := strconv.ParseUint(idParam, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   "Geçersiz company id",
        })
        return
    }

    if err := h.service.DeleteCompany(uint(id)); err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "success": false,
            "error":   "Company bulunamadı",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Company silindi",
    })
}
