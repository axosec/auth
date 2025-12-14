package api

import (
	"net/http"

	"github.com/axosec/auth/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetSelfHandler(c *gin.Context) {

	user := c.MustGet("user").(dto.User)

	c.JSON(http.StatusOK, user)
}
