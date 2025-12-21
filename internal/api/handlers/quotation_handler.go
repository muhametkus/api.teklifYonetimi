package handlers

import (
	"bytes"
	"net/http"
	"path/filepath"
	"strconv"

	"html/template"

	"api.teklifYonetimi/internal/api/response"
	"api.teklifYonetimi/internal/dto"
	"api.teklifYonetimi/internal/models"
	"api.teklifYonetimi/internal/repository"
	"api.teklifYonetimi/internal/services"
	"api.teklifYonetimi/internal/utils"

	"github.com/gin-gonic/gin"
)

type QuotationHandler struct {
	service *services.QuotationService
}

func NewQuotationHandler() *QuotationHandler {
	repo := repository.NewQuotationRepository()
	service := services.NewQuotationService(repo)

	return &QuotationHandler{
		service: service,
	}
}

//
// CREATE QUOTATION
// POST /quotations
//
// CreateQuotation godoc
// @Summary      Yeni Teklif Oluştur
// @Description  Yeni bir teklif oluşturur ve kaydeder
// @Tags         Quotations
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        quotation body dto.CreateQuotationRequest true "Teklif Bilgileri"
// @Success      201  {object} response.APIResponse
// @Failure      400  {object} response.APIResponse
// @Router       /quotations [post]
func (h *QuotationHandler) CreateQuotation(c *gin.Context) {

	var req struct {
		Title       string                 `json:"title" binding:"required"`
		Customer    string                 `json:"customer" binding:"required"`
		Description string                 `json:"description"`
		Items       []models.QuotationItem `json:"items" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		status, res := response.Error(
			http.StatusBadRequest,
			"Geçersiz istek",
			"VALIDATION_ERROR",
		)
		c.JSON(status, res)
		return
	}

	companyID := c.MustGet("company_id").(uint)
	userID := c.MustGet("user_id").(uint)

	quotation, err := h.service.CreateQuotation(
		companyID,
		userID,
		req.Title,
		req.Customer,
		req.Description,
		req.Items,
	)
	if err != nil {
		status, res := response.Error(
			http.StatusBadRequest,
			err.Error(),
			"CREATE_QUOTATION_FAILED",
		)
		c.JSON(status, res)
		return
	}

	status, res := response.Created(
		"Teklif başarıyla oluşturuldu",
		quotation,
	)
	c.JSON(status, res)
}

//
// LIST QUOTATIONS (PAGINATION + FILTER + ROLE)
// GET /quotations
//
// GetQuotations godoc
// @Summary      Teklifleri Listele
// @Description  Filtreleme ve sayfalama ile teklifleri listeler
// @Tags         Quotations
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        page query int false "Sayfa Numarası"
// @Param        limit query int false "Sayfa Başına Kayıt"
// @Param        status query string false "Durum Filtresi"
// @Param        customer query string false "Müşteri Filtresi"
// @Success      200  {object} response.APIResponse
// @Failure      500  {object} response.APIResponse
// @Router       /quotations [get]
func (h *QuotationHandler) GetQuotations(c *gin.Context) {

	companyID := c.MustGet("company_id").(uint)
	userID := c.MustGet("user_id").(uint)
	role := c.MustGet("role").(string)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	statusFilter := c.Query("status")
	customerFilter := c.Query("customer")

	quotations, total, err := h.service.GetFilteredPaginatedQuotations(
		companyID,
		userID,
		role,
		statusFilter,
		customerFilter,
		page,
		limit,
	)
	if err != nil {
		status, res := response.Error(
			http.StatusInternalServerError,
			"Teklifler alınamadı",
			"QUOTATION_LIST_FAILED",
		)
		c.JSON(status, res)
		return
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	statusCode, res := response.Success(
		"Teklifler başarıyla getirildi",
		quotations,
		gin.H{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"totalPages": totalPages,
		},
	)

	c.JSON(statusCode, res)
}

//
// GET QUOTATION PDF
// GET /quotations/:id/pdf
//
func (h *QuotationHandler) GetQuotationPDF(c *gin.Context) {

	// Param
	idParam := c.Param("id")
	quotationID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		status, res := response.Error(
			http.StatusBadRequest,
			"Geçersiz teklif id",
			"INVALID_ID",
		)
		c.JSON(status, res)
		return
	}

	companyID := c.MustGet("company_id").(uint)
	userID := c.MustGet("user_id").(uint)
	role := c.MustGet("role").(string)

	quotation, err := h.service.GetQuotationForPDF(
		companyID,
		uint(quotationID),
	)
	if err != nil {
		status, res := response.Error(
			http.StatusNotFound,
			"Teklif bulunamadı",
			"QUOTATION_NOT_FOUND",
		)
		c.JSON(status, res)
		return
	}

	// OWNERSHIP CHECK (USER)
	if role == "USER" && quotation.CreatedBy != userID {
		status, res := response.Error(
			http.StatusForbidden,
			"Bu teklife erişim yetkiniz yok",
			"FORBIDDEN",
		)
		c.JSON(status, res)
		return
	}

	// CACHE CHECK
	cacheKey := "quotation_" + idParam + "_" + string(quotation.Status)

	if cachedPDF, ok := utils.GetPDF(cacheKey); ok {
		c.Header("Content-Type", "application/pdf")
		c.Header("Content-Disposition", "inline; filename=quotation.pdf")
		c.Data(http.StatusOK, "application/pdf", cachedPDF)
		return
	}

	// ViewModel
	var logoBase64 string
	if quotation.Company.Logo != nil {
		logoBase64, _ = utils.ImageToBase64(*quotation.Company.Logo)
	}

	view := dto.QuotationPDFView{
		Title:       quotation.Title,
		Customer:    quotation.Customer,
		Status:      string(quotation.Status),
		Total:       quotation.Total,
		Items:       quotation.Items,
		CompanyName: quotation.Company.Name,
		CompanyLogo: logoBase64,
	}

	// HTML render
	templatePath, _ := filepath.Abs("internal/templates/quotation.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		status, res := response.Error(
			http.StatusInternalServerError,
			"Template okunamadı",
			"TEMPLATE_ERROR",
		)
		c.JSON(status, res)
		return
	}

	var htmlBuffer bytes.Buffer
	if err := tmpl.Execute(&htmlBuffer, view); err != nil {
		status, res := response.Error(
			http.StatusInternalServerError,
			"Template render edilemedi",
			"TEMPLATE_RENDER_ERROR",
		)
		c.JSON(status, res)
		return
	}

	// HTML → PDF
	pdfBytes, err := utils.GeneratePDFFromHTMLBytes(htmlBuffer.String())
	if err != nil {
		status, res := response.Error(
			http.StatusInternalServerError,
			"PDF oluşturulamadı",
			"PDF_GENERATE_ERROR",
		)
		c.JSON(status, res)
		return
	}

	// CACHE WRITE
	utils.SetPDF(cacheKey, pdfBytes)

	// RESPONSE
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "inline; filename=quotation.pdf")
	c.Data(http.StatusOK, "application/pdf", pdfBytes)
}

//
// UPDATE STATUS
// PUT /quotations/:id/status
//
func (h *QuotationHandler) UpdateQuotationStatus(c *gin.Context) {

	idParam := c.Param("id")
	quotationID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		status, res := response.Error(
			http.StatusBadRequest,
			"Geçersiz ID",
			"INVALID_ID",
		)
		c.JSON(status, res)
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		status, res := response.Error(
			http.StatusBadRequest,
			"Status zorunludur",
			"VALIDATION_ERROR",
		)
		c.JSON(status, res)
		return
	}

	companyID := c.MustGet("company_id").(uint)

	if err := h.service.UpdateQuotationStatus(companyID, uint(quotationID), req.Status); err != nil {
		status, res := response.Error(
			http.StatusInternalServerError,
			"Durum güncellenemedi",
			"UPDATE_FAILED",
		)
		c.JSON(status, res)
		return
	}

	status, res := response.Success(
		"Durum güncellendi",
		nil,
		nil,
	)
	c.JSON(status, res)
}

// GetQuotationByID
// GET /quotations/:id
// GetQuotationByID godoc
// @Summary      Teklif Detayı
// @Description  Teklif detayını getirir
// @Tags         Quotations
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Teklif ID"
// @Success      200  {object} response.APIResponse
// @Failure      404  {object} response.APIResponse
// @Router       /quotations/{id} [get]
func (h *QuotationHandler) GetQuotationByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		status, res := response.Error(http.StatusBadRequest, "Geçersiz ID", "INVALID_ID")
		c.JSON(status, res)
		return
	}

	companyID := c.MustGet("company_id").(uint)

	quotation, err := h.service.GetQuotationByID(uint(id), companyID)
	if err != nil {
		status, res := response.Error(http.StatusNotFound, "Teklif bulunamadı", "NOT_FOUND")
		c.JSON(status, res)
		return
	}

	status, res := response.Success("Teklif detayı", quotation, nil)
	c.JSON(status, res)
}

// UpdateQuotation
// PUT /quotations/:id
// UpdateQuotation godoc
// @Summary      Teklif Güncelle
// @Description  Teklif içeriğini günceller
// @Tags         Quotations
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Teklif ID"
// @Param        quotation body dto.UpdateQuotationRequest true "Güncelleme Bilgileri"
// @Success      200  {object} response.APIResponse
// @Failure      400  {object} response.APIResponse
// @Router       /quotations/{id} [put]
func (h *QuotationHandler) UpdateQuotation(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		status, res := response.Error(http.StatusBadRequest, "Geçersiz ID", "INVALID_ID")
		c.JSON(status, res)
		return
	}

	var req dto.UpdateQuotationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		status, res := response.Error(http.StatusBadRequest, err.Error(), "VALIDATION_ERROR")
		c.JSON(status, res)
		return
	}

	companyID := c.MustGet("company_id").(uint)
	userID := c.MustGet("user_id").(uint)
	role := c.MustGet("role").(string)

	// Convert DTO items to Model items
	var items []models.QuotationItem
	for _, itemReq := range req.Items {
		items = append(items, models.QuotationItem{
			ItemName:  itemReq.ItemName,
			Quantity:  itemReq.Quantity,
			UnitPrice: itemReq.UnitPrice,
		})
	}

	quotation, err := h.service.UpdateQuotation(
		uint(id), companyID, userID, role,
		req.Title, req.Customer, req.Description, items,
	)
	if err != nil {
		status, res := response.Error(http.StatusInternalServerError, err.Error(), "UPDATE_FAILED")
		c.JSON(status, res)
		return
	}

	status, res := response.Success("Teklif güncellendi", quotation, nil)
	c.JSON(status, res)
}

// DeleteQuotation
// DELETE /quotations/:id
// DeleteQuotation godoc
// @Summary      Teklif Sil
// @Description  Teklifi siler
// @Tags         Quotations
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Teklif ID"
// @Success      200  {object} response.APIResponse
// @Failure      404  {object} response.APIResponse
// @Router       /quotations/{id} [delete]
func (h *QuotationHandler) DeleteQuotation(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		status, res := response.Error(http.StatusBadRequest, "Geçersiz ID", "INVALID_ID")
		c.JSON(status, res)
		return
	}

	companyID := c.MustGet("company_id").(uint)
	userID := c.MustGet("user_id").(uint)
	role := c.MustGet("role").(string)

	if err := h.service.DeleteQuotation(uint(id), companyID, userID, role); err != nil {
		status, res := response.Error(http.StatusInternalServerError, err.Error(), "DELETE_FAILED")
		c.JSON(status, res)
		return
	}

	status, res := response.Success("Teklif silindi", nil, nil)
	c.JSON(status, res)
}
