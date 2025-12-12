package routes

import (
    "api.teklifYonetimi/internal/api/handlers"

    "github.com/gin-gonic/gin"
)

func RegisterCompanyRoutes(r *gin.Engine) {
    companyHandler := handlers.NewCompanyHandler()

    companies := r.Group("/companies")
    {
        companies.POST("", companyHandler.CreateCompany)
        companies.GET("", companyHandler.GetCompanies)
		companies.GET("/:id", companyHandler.GetCompanyByID)
		companies.PUT("/:id", companyHandler.UpdateCompany)
    }
}