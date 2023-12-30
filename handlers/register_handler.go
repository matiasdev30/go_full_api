package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/matiasdev30/go_api/db"
	"github.com/matiasdev30/go_api/models"
	"github.com/matiasdev30/go_api/service"
	"github.com/matiasdev30/go_api/utils"
)

// @BasePath /api/v1

// @Summary Register user
// @Description Register new user
// @Tags Register
// @Accept json
// @Produce json
// @Param request body UserResponse true "Request body"
// @Success 200 {object} UserResponse
// @Router /auth/register [post]
func RegisterHanlder(ctx *gin.Context) {

	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		utils.SendErro(ctx, http.StatusBadRequest, "cannot bind json "+err.Error())
		return
	}

	if len(user.Password) < 8 {
		utils.SendErro(ctx, http.StatusBadRequest, "password need to contains 8 word or more")
		return
	}

	if user.PhoneNumber != 0 && len(fmt.Sprintf("%x", user.PhoneNumber)) < 9 {
		utils.SendErro(ctx, http.StatusBadRequest, "invalid phone number")
		return
	}

	if !strings.Contains(user.Email, "@") {
		utils.SendErro(ctx, http.StatusBadRequest, "invalid email")
		return
	}

	userRgisted := models.User{}

	if db.Database.Find(&userRgisted, "email = ?", user.Email); userRgisted.Email == user.Email {
		utils.SendErro(ctx, http.StatusBadRequest, "email just contains")
		return
	}

	user.Password = service.SHA256Encoder(user.Password)

	if err := db.Database.Create(&user).Error; err != nil {
		utils.SendErro(ctx, http.StatusInternalServerError, "error created user"+err.Error())
		return
	}

	utils.SendSucess(ctx, http.StatusOK, user, "data")

}
