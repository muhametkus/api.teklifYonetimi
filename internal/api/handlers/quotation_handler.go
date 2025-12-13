package handlers

import (
	"net/http"

	"api.teklifYonetimi/internal/dto"
	"api.teklifYonetimi/internal/models"
	"api.teklifYonetimi/internal/repository"
	"api.teklifYonetimi/internal/services"

	"github.com/gin-gonic/gin"
	"strconv"
	"bytes"
	"html/template"
	"path/filepath"
	
	
	
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

// CreateQuotation
// POST /quotations
func (h *QuotationHandler) CreateQuotation(c *gin.Context) {
	var req dto.CreateQuotationRequest

	// 1️⃣ JSON -> DTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// 2️⃣ JWT'den company_id al
	companyIDValue, exists := c.Get("company_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Company bilgisi bulunamadı",
		})
		return
	}

	companyID := companyIDValue.(uint)

	// 3️⃣ DTO -> Model (Item'lar)
	var items []models.QuotationItem
	for _, item := range req.Items {
		items = append(items, models.QuotationItem{
			ItemName:  item.ItemName,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		})
	}

	// 4️⃣ Service çağır
	quotation, err := h.service.CreateQuotation(
		companyID,
		req.Title,
		req.Customer,
		req.Description,
		items,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Teklif oluşturulamadı",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    quotation,
	})
}

// GetQuotations
// GET /quotations
func (h *QuotationHandler) GetQuotations(c *gin.Context) {

	companyIDValue, exists := c.Get("company_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Company bilgisi bulunamadı",
		})
		return
	}

	companyID := companyIDValue.(uint)

	quotations, err := h.service.GetCompanyQuotations(companyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Teklifler alınamadı",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    quotations,
	})
}

// UpdateQuotationStatus
// PUT /quotations/:id/status
func (h *QuotationHandler) UpdateQuotationStatus(c *gin.Context) {

	idParam := c.Param("id")
	quotationID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Geçersiz teklif id",
		})
		return
	}

	var req dto.UpdateQuotationStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	companyID := c.MustGet("company_id").(uint)

	err = h.service.UpdateQuotationStatus(
		companyID,
		uint(quotationID),
		req.Status,
	)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Teklif durumu güncellendi",
	})
}

// GetQuotations
// GET /quotations?page=1&limit=10
func (h *QuotationHandler) GetQuotations(c *gin.Context) {

	companyID := c.MustGet("company_id").(uint)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	quotations, total, err := h.service.GetCompanyQuotationsPaginated(
		companyID,
		page,
		limit,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Teklifler alınamadı",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    quotations,
		"meta": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

// GetQuotationPDF
// GET /quotations/:id/pdf
func (h *QuotationHandler) GetQuotationPDF(c *gin.Context) {

	// 1️⃣ Parametre
	idParam := c.Param("id")
	quotationID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Geçersiz teklif id",
		})
		return
	}

	// 2️⃣ JWT'den company_id
	companyID := c.MustGet("company_id").(uint)

	// 3️⃣ Teklifi al
	quotation, err := h.service.GetQuotationForPDF(
		companyID,
		uint(quotationID),
	)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Teklif bulunamadı",
		})
		return
	}

	// 4️⃣ HTML template render
	templatePath, _ := filepath.Abs("internal/templates/quotation.html")

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Template okunamadı",
		})
		return
	}

	var htmlBuffer bytes.Buffer
	err = tmpl.Execute(&htmlBuffer, quotation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Template render edilemedi",
		})
		return
	}

	// 5️⃣ HTML → PDF
	pdfPath := "quotation_" + idParam + ".pdf"

	err = utils.GeneratePDFFromHTML(
		htmlBuffer.String(),
		pdfPath,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "PDF oluşturulamadı",
		})
		return
	}

	// 6️⃣ PDF'i response olarak gönder
	c.File(pdfPath)
}
