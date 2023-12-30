package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/matiasdev30/go_api/models"
)

func SendSucess(ctx *gin.Context, code int, data any, op string) {

	ctx.JSON(
		code, gin.H{
			op: data,
		},
	)
}

func SendErro(ctx *gin.Context, code int, data any) {
	ctx.JSON(
		code, gin.H{
			"error": data,
		},
	)
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateTaskResponse struct {
	Message string      `json:"message"`
	Data    models.Task `json:"data"`
}
