package api

import (
	"github.com/axosec/auth/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService *service.AuthService
}

func NewHandler(authService *service.AuthService) *Handler {
	return &Handler{
		authService: authService,
	}
}

// @Summary      Helthcheck
// @Description  returns ok if api up
// @Success      200
// @Router       /health [get]
func (h *Handler) Helth(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}

func (h *Handler) RegisterRouters(e *gin.Engine) {
	v1 := e.Group("/v1")
	{
		v1.GET("/health", h.Helth)

		auth := v1.Group("/auth")
		{
			auth.POST("/register", h.RegisterHandler)
			auth.POST("/login/init", h.InitLoginHandeler)
			auth.POST("/login", h.LoginHandeler)

		}
	}
}
