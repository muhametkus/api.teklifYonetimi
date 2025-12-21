package routes

import (
	"api.teklifYonetimi/internal/api/handlers"
	"api.teklifYonetimi/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterQuotationRoutes(r *gin.Engine) {
	quotationHandler := handlers.NewQuotationHandler()

	quotations := r.Group("/quotations")
	quotations.Use(middleware.JWTAuthMiddleware())
	{
		// Create & List
		quotations.POST("", quotationHandler.CreateQuotation)
		quotations.GET("", quotationHandler.GetQuotations)

		// Detail & PDF
		quotations.GET("/:id", quotationHandler.GetQuotationByID)
		quotations.GET("/:id/pdf", quotationHandler.GetQuotationPDF)

		// Update & Delete
		quotations.PUT("/:id", quotationHandler.UpdateQuotation)
		quotations.DELETE("/:id", quotationHandler.DeleteQuotation)

		// Status Update (ADMIN only)
		quotations.PUT(
			"/:id/status",
			middleware.RequireRole("ADMIN"),
			quotationHandler.UpdateQuotationStatus,
		)


	}
}


