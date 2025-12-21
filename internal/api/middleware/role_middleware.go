package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequireRole
// Örn: RequireRole("ADMIN")
func RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {

		roleValue, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"success": false,
				"error":   "Rol bilgisi bulunamadı",
			})
			return
		}

		role, ok := roleValue.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"success": false,
				"error":   "Rol formatı hatalı",
			})
			return
		}

		// SUPER_ADMIN her yere erişebilir
		if role == "SUPER_ADMIN" {
			c.Next()
			return
		}

		if role != requiredRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"success": false,
				"error":   "Bu işlem için yetkiniz yok",
			})
			return
		}

		c.Next()
	}
}
