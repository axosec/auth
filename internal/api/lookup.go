package api

import (
	"net/http"

	"github.com/axosec/auth/internal/dto"
	"github.com/axosec/auth/internal/service"
	"github.com/gin-gonic/gin"
)

// @Summary      Lookup user by email hash
// @Description  returns a single user matching the email hash
// @Success      200 {object} dto.UserLookup
// @Param		 request body dto.LookupUserRequest true "Lookup payload"
// @Router       /user/lookup [post]
func (h *Handler) LookupUserHandler(c *gin.Context) {
	var req dto.LookupUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.LookupUser(c, req.EmailHash)
	if err != nil {
		if err == service.ErrUserNotFound {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary      Lookup users by IDs
// @Description  returns users matching the provided IDs
// @Success      200 {array} dto.UserLookup
// @Param		 request body dto.LookupUsersRequest true "Lookup payload"
// @Router       /users/lookup [post]
func (h *Handler) LookupUsersHandler(c *gin.Context) {
	var req dto.LookupUsersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := h.userService.LookupUsers(c, req.IDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, users)
}
