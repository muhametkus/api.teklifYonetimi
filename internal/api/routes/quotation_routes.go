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
		quotations.POST("", quotationHandler.CreateQuotation)
		quotations.GET("", quotationHandler.GetQuotations)
		quotations.PUT(
	"/:id/status",
	middleware.RequireRole("ADMIN"),
	quotationHandler.UpdateQuotationStatus,
)
quotations.GET(
	"/:id/pdf",
	quotationHandler.GetQuotationPDF,
)


	}
}


