package handlers

import (
	"net/http"

	"api.teklifYonetimi/internal/api/response"
	"api.teklifYonetimi/internal/dto"
	"api.teklifYonetimi/internal/repository"
	"api.teklifYonetimi/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler() *AuthHandler {
	userRepo := repository.NewUserRepository()
	authService := services.NewAuthService(userRepo)

	return &AuthHandler{
		authService: authService,
	}
}

// Login
// POST /auth/login
// Login godoc
// @Summary      Kullanıcı Girişi
// @Description  Email ve şifre ile giriş yapar, JWT token döner
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        login body dto.LoginRequest true "Giriş Bilgileri"
// @Success      200  {object} response.APIResponse
// @Failure      400  {object} response.APIResponse
// @Failure      401  {object} response.APIResponse
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		status, res := response.Error(http.StatusBadRequest, err.Error(), "VALIDATION_ERROR")
		c.JSON(status, res)
		return
	}

	token, user, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		status, res := response.Error(http.StatusUnauthorized, err.Error(), "LOGIN_FAILED")
		c.JSON(status, res)
		return
	}

	status, res := response.Success("Giriş başarılı", gin.H{"token": token, "user": user}, nil)
	c.JSON(status, res)
}
