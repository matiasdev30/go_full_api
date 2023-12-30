package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matiasdev30/go_api/service"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const Bearer_scheme = "Bearer "
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			ctx.Abort()
			return
		}

		token := header[len(Bearer_scheme):]

		if isvalid, _ := service.VerifyToken(token); !isvalid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
