package handlers

import (
    "net/http"
    "strconv"

    "api.teklifYonetimi/internal/api/response"
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
// CreateCompany godoc
// @Summary      Yeni Şirket Oluştur
// @Description  Yeni bir şirket kaydı oluşturur
// @Tags         Companies
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        company body dto.CreateCompanyRequest true "Şirket Bilgileri"
// @Success      201  {object} response.APIResponse
// @Failure      400  {object} response.APIResponse
// @Failure      500  {object} response.APIResponse
// @Router       /companies [post]
func (h *CompanyHandler) CreateCompany(c *gin.Context) {
    var req dto.CreateCompanyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		status, res := response.Error(http.StatusBadRequest, err.Error(), "VALIDATION_ERROR")
		c.JSON(status, res)
		return
	}

	company, err := h.service.CreateCompany(req.Name, req.Logo)
	if err != nil {
		status, res := response.Error(http.StatusInternalServerError, "Company oluşturulamadı", "CREATE_FAILED")
		c.JSON(status, res)
		return
	}

	status, res := response.Created("Şirket başarıyla oluşturuldu", company)
	c.JSON(status, res)
}

// GetCompanies
// GET /companies
// GetCompanies godoc
// @Summary      Şirketleri Listele
// @Description  Tüm şirketleri listeler
// @Tags         Companies
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object} response.APIResponse
// @Failure      500  {object} response.APIResponse
// @Router       /companies [get]
func (h *CompanyHandler) GetCompanies(c *gin.Context) {
	companies, err := h.service.GetAllCompanies()
	if err != nil {
		status, res := response.Error(http.StatusInternalServerError, "Company listesi alınamadı", "LIST_FAILED")
		c.JSON(status, res)
		return
	}

	status, res := response.Success("Şirketler listelendi", companies, nil)
	c.JSON(status, res)
}

// GetCompanyByID
// GET /companies/:id
// GetCompanyByID godoc
// @Summary      Şirket Detayı
// @Description  ID'ye göre şirket detayını getirir
// @Tags         Companies
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Şirket ID"
// @Success      200  {object} response.APIResponse
// @Failure      404  {object} response.APIResponse
// @Router       /companies/{id} [get]
func (h *CompanyHandler) GetCompanyByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		status, res := response.Error(http.StatusBadRequest, "Geçersiz company id", "INVALID_ID")
		c.JSON(status, res)
		return
	}

	company, err := h.service.GetCompanyByID(uint(id))
	if err != nil {
		status, res := response.Error(http.StatusNotFound, "Company bulunamadı", "NOT_FOUND")
		c.JSON(status, res)
		return
	}

	status, res := response.Success("Şirket detayı getirildi", company, nil)
	c.JSON(status, res)
}

// UpdateCompany
// PUT /companies/:id
// UpdateCompany godoc
// @Summary      Şirket Güncelle
// @Description  Şirket bilgilerini günceller
// @Tags         Companies
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Şirket ID"
// @Param        company body dto.UpdateCompanyRequest true "Güncelleme Bilgileri"
// @Success      200  {object} response.APIResponse
// @Failure      400  {object} response.APIResponse
// @Router       /companies/{id} [put]
func (h *CompanyHandler) UpdateCompany(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		status, res := response.Error(http.StatusBadRequest, "Geçersiz company id", "INVALID_ID")
		c.JSON(status, res)
		return
	}

	var req dto.UpdateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		status, res := response.Error(http.StatusBadRequest, err.Error(), "VALIDATION_ERROR")
		c.JSON(status, res)
		return
	}

	company, err := h.service.UpdateCompany(
		uint(id),
		req.Name,
		req.Logo,
		req.Subscription,
	)

	if err != nil {
		status, res := response.Error(http.StatusNotFound, "Company bulunamadı veya güncellenemedi", "UPDATE_FAILED")
		c.JSON(status, res)
		return
	}

	status, res := response.Success("Şirket güncellendi", company, nil)
	c.JSON(status, res)
}


// DeleteCompany
// DELETE /companies/:id
// DeleteCompany godoc
// @Summary      Şirket Sil
// @Description  Şirketi siler
// @Tags         Companies
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Şirket ID"
// @Success      200  {object} response.APIResponse
// @Failure      404  {object} response.APIResponse
// @Router       /companies/{id} [delete]
func (h *CompanyHandler) DeleteCompany(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		status, res := response.Error(http.StatusBadRequest, "Geçersiz company id", "INVALID_ID")
		c.JSON(status, res)
		return
	}

	if err := h.service.DeleteCompany(uint(id)); err != nil {
		status, res := response.Error(http.StatusNotFound, "Company bulunamadı", "DELETE_FAILED")
		c.JSON(status, res)
		return
	}

	status, res := response.Success("Şirket silindi", nil, nil)
	c.JSON(status, res)
}
