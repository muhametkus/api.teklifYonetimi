package middleware

import (
	"net/http"
	"strings"

	"api.teklifYonetimi/internal/repository"
	"api.teklifYonetimi/internal/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 1) Authorization header al
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Authorization header eksik",
			})
			return
		}

		// 2) Bearer token formatını kontrol et
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Authorization formatı hatalı",
			})
			return
		}

		tokenString := parts[1]

		// 3) Token parse et
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Geçersiz token",
			})
			return
		}

		// 4) DB'den user'ı bul (company_id için)
		userRepo := repository.NewUserRepository()
		user, err := userRepo.FindByID(claims.UserID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Kullanıcı bulunamadı",
			})
			return
		}

		if user.CompanyID == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Kullanıcının company bilgisi yok",
			})
			return
		}

		// 5) Context'e user bilgilerini koy
		c.Set("user_id", user.ID)
		c.Set("role", string(user.Role))
		c.Set("company_id", *user.CompanyID)

		c.Next()
	}
}
