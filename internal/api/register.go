package api

import (
	"net/http"

	"github.com/axosec/auth/internal/dto"
	"github.com/axosec/auth/internal/service"
	"github.com/gin-gonic/gin"
)

// @Summary      Register user
// @Description  registers new user
// @Success      201
// @Param user body dto.RegisterRequest true "User payload"
// @Router       /auth/register [post]
func (h *Handler) RegisterHandler(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	if err := h.authService.RegisterUser(c.Request.Context(), req); err != nil {
		if err == service.ErrInvalidKeyLength || err == service.ErrUserAlreadyExists {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}
