package api

import (
	"net/http"

	"github.com/axosec/auth/internal/dto"
	"github.com/gin-gonic/gin"
)

// @Summary      Init Login
// @Description  initializes login by obtaining users salt
// @Success      200
// @Param user body dto.InitLoginRequest true "Login request"
// @Router       /auth/login/init [post]
func (h *Handler) InitLoginHandeler(c *gin.Context) {
	var req dto.InitLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	res, err := h.authService.InitLogin(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not initialize login"})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary      Login
// @Description  logs user in
// @Success      200
// @Param user body dto.LoginRequest true "Login request"
// @Router       /auth/login [post]
func (h *Handler) LoginHandeler(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	user, token, err := h.authService.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not login"})
		return
	}
	c.SetCookie("auth_token", token, 3*24*60*60, "/", "", false, true)
	c.JSON(http.StatusCreated, user)
}

// @Summary      Logout
// @Description  logs user out
// @Success      200
// @Router       /user/logout [post]
func (h *Handler) LogoutHandler(c *gin.Context) {
	c.SetCookie("auth_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User logged out in successfully",
	})
}
