package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/matiasdev30/go_api/db"
	"github.com/matiasdev30/go_api/models"
	"github.com/matiasdev30/go_api/service"
	"github.com/matiasdev30/go_api/utils"
)

// @BasePath /api/v1

// @Summary Login
// @Description Authentication user
// @Tags Login
// @Accept json
// @Produce json
// @Param request body models.Login true "Request body"
// @Success 200 {object} LoginResponse
// @Router /auth/login [post]
func LoginHandler(ctx *gin.Context) {

	var data = models.Login{}

	err := ctx.ShouldBindJSON(&data)

	if err != nil {
		utils.SendErro(ctx, http.StatusBadRequest, "cannot bind json "+err.Error())
		return
	}

	if data.Email == "" || !strings.Contains(data.Email, "@") {
		utils.SendErro(ctx, http.StatusBadRequest, "invalid email")
		return
	}

	var user models.User
	if err := db.Database.Where("email=?", data.Email).First(&user).Error; err != nil {
		utils.SendErro(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if user.Password != service.SHA256Encoder(data.Password) {
		utils.SendErro(ctx, http.StatusBadRequest, "invalid password")
		return
	}

	token, err := service.GenereteToken(&user)

	if err != nil {
		utils.SendErro(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSucess(ctx, http.StatusAccepted, map[string]any{"user": user, "token": token}, "data")

}

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type UserResponse struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber int    `json:"phoneNumber"`
}
