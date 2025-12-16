package api

import (
	"github.com/axosec/auth/internal/service"
	"github.com/axosec/core/crypto/token"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	jwt         *token.JWTManager
	authService *service.AuthService
	userService *service.UserService
}

func NewHandler(jwt *token.JWTManager, authService *service.AuthService, userService *service.UserService) *Handler {
	return &Handler{
		jwt:         jwt,
		authService: authService,
		userService: userService,
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
		authenticated := v1.Group("")
		authenticated.Use(h.AuthenticatedMiddleware())
		{
			authenticated.GET("/user/self", h.GetSelfHandler)
			authenticated.POST("/user/logout", h.LogoutHandler)
		}
	}
}
