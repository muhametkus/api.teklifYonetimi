package handlers

import (
	"net/http"
	"strconv"

	"api.teklifYonetimi/internal/api/response"
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
// CreateUser godoc
// @Summary      Yeni Kullanıcı Oluştur
// @Description  Yeni bir kullanıcı kaydı oluşturur
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        user body dto.CreateUserRequest true "Kullanıcı Bilgileri"
// @Success      201  {object} response.APIResponse
// @Failure      400  {object} response.APIResponse
// @Router       /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		status, res := response.Error(http.StatusBadRequest, err.Error(), "VALIDATION_ERROR")
		c.JSON(status, res)
		return
	}

	// Context'ten role ve company_id al (Middleware varsa)
	role, _ := c.Get("role")
	companyID, _ := c.Get("company_id")

	// Eğer istek yapan ADMIN ise, oluşturulan user'ı kendi şirketine bağla
	if role == "ADMIN" && companyID != nil {
		cID := companyID.(uint)
		req.CompanyID = &cID
	}

	user, err := h.service.CreateUser(
		req.Name,
		req.Email,
		req.Password,
		req.Role,
		req.CompanyID,
	)
	if err != nil {
		status, res := response.Error(http.StatusBadRequest, err.Error(), "CREATE_FAILED")
		c.JSON(status, res)
		return
	}

	// Password'ü response'tan çıkaralım
	user.Password = ""

	status, res := response.Created("Kullanıcı başarıyla oluşturuldu", user)
	c.JSON(status, res)
}

// GetUsers
// GET /users
// GetUsers godoc
// @Summary      Kullanıcıları Listele
// @Description  Tüm kullanıcıları listeler
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object} response.APIResponse
// @Failure      500  {object} response.APIResponse
// @Router       /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		status, res := response.Error(http.StatusInternalServerError, "User listesi alınamadı", "LIST_FAILED")
		c.JSON(status, res)
		return
	}

	// Password'leri response'tan çıkar
	for i := range users {
		users[i].Password = ""
	}

	status, res := response.Success("Kullanıcılar listelendi", users, nil)
	c.JSON(status, res)
}

// GetUserByID
// GET /users/:id
// GetUserByID godoc
// @Summary      Kullanıcı Detayı
// @Description  ID'ye göre kullanıcı detayını getirir
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Kullanıcı ID"
// @Success      200  {object} response.APIResponse
// @Failure      404  {object} response.APIResponse
// @Router       /users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		status, res := response.Error(http.StatusBadRequest, "Geçersiz ID", "INVALID_ID")
		c.JSON(status, res)
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		status, res := response.Error(http.StatusNotFound, "Kullanıcı bulunamadı", "NOT_FOUND")
		c.JSON(status, res)
		return
	}

	user.Password = ""
	status, res := response.Success("Kullanıcı detayı", user, nil)
	c.JSON(status, res)
}

// UpdateUser
// PUT /users/:id
// UpdateUser godoc
// @Summary      Kullanıcı Güncelle
// @Description  Kullanıcı bilgilerini günceller
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Kullanıcı ID"
// @Param        user body dto.UpdateUserRequest true "Güncelleme Bilgileri"
// @Success      200  {object} response.APIResponse
// @Failure      400  {object} response.APIResponse
// @Router       /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		status, res := response.Error(http.StatusBadRequest, "Geçersiz ID", "INVALID_ID")
		c.JSON(status, res)
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		status, res := response.Error(http.StatusBadRequest, err.Error(), "VALIDATION_ERROR")
		c.JSON(status, res)
		return
	}

	user, err := h.service.UpdateUser(uint(id), req.Name, req.Email, req.Password, req.Role)
	if err != nil {
		status, res := response.Error(http.StatusInternalServerError, "Güncelleme başarısız", "UPDATE_FAILED")
		c.JSON(status, res)
		return
	}

	user.Password = ""
	status, res := response.Success("Kullanıcı güncellendi", user, nil)
	c.JSON(status, res)
}

// DeleteUser
// DELETE /users/:id
// DeleteUser godoc
// @Summary      Kullanıcı Sil
// @Description  Kullanıcıyı siler
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Kullanıcı ID"
// @Success      200  {object} response.APIResponse
// @Failure      404  {object} response.APIResponse
// @Router       /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		status, res := response.Error(http.StatusBadRequest, "Geçersiz ID", "INVALID_ID")
		c.JSON(status, res)
		return
	}

	if err := h.service.DeleteUser(uint(id)); err != nil {
		status, res := response.Error(http.StatusInternalServerError, "Silme başarısız", "DELETE_FAILED")
		c.JSON(status, res)
		return
	}

	status, res := response.Success("Kullanıcı silindi", nil, nil)
	c.JSON(status, res)
}
