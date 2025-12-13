package handlers

import (
	"net/http"

	"api.teklifYonetimi/internal/dto"
	"api.teklifYonetimi/internal/repository"
	"api.teklifYonetimi/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler() *UserHandler {
	repo := repository.NewUserRepository()
	service := services.NewUserService(repo)

	return &UserHandler{
		service: service,
	}
}

// CreateUser
// POST /users
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	user, err := h.service.CreateUser(
		req.Name,
		req.Email,
		req.Password,
		req.Role,
		req.CompanyID,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Password'ü response'tan çıkaralım
	user.Password = ""

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    user,
	})
}

// GetUsers
// GET /users
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "User listesi alınamadı",
		})
		return
	}

	// Password'leri response'tan çıkar
	for i := range users {
		users[i].Password = ""
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
	})
}
