package middlewares

import (
	"github.com/duynguyenbui/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		if token != "valid-token" {
			response.ErrorResponse(ctx, response.ErrInvalidToken, "Invalid token")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
