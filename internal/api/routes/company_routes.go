package routes

import (
    "api.teklifYonetimi/internal/api/handlers"

    "github.com/gin-gonic/gin"

    "api.teklifYonetimi/internal/api/middleware"

)

func RegisterCompanyRoutes(r *gin.Engine) {
	companyHandler := handlers.NewCompanyHandler()

	companies := r.Group("/companies")
	companies.Use(middleware.JWTAuthMiddleware())
	{
		// HERKES (USER + ADMIN)
		companies.GET("", companyHandler.GetCompanies)
		companies.GET("/:id", companyHandler.GetCompanyByID)

		// SADECE ADMIN
		companies.POST(
			"",
			middleware.RequireRole("ADMIN"),
			companyHandler.CreateCompany,
		)

		companies.PUT(
			"/:id",
			middleware.RequireRole("ADMIN"),
			companyHandler.UpdateCompany,
		)

		companies.DELETE(
			"/:id",
			middleware.RequireRole("ADMIN"),
			companyHandler.DeleteCompany,
		)
	}
}

