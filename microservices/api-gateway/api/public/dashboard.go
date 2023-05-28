package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Dashboard godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func Dashboard(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, `{"dasboard":"cool"}`)
}
