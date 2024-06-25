package initialize

import (
	"fmt"

	c "github.com/duynguyenbui/go-ecommerce-backend-api/internal/controller"
	"github.com/duynguyenbui/go-ecommerce-backend-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Before --> AA")
		ctx.Next()
		fmt.Println("After --> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Before --> BB")
		ctx.Next()
		fmt.Println("After --> BB")
	}
}

func CC(ctx *gin.Context) {
	fmt.Println("Before --> CC")
	ctx.Next()
	fmt.Println("After --> CC")
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	// use the middlewares
	r.Use(middlewares.AuthenMiddleware(), BB(), CC)

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", c.NewPongController().Pong)
		v1.GET("/user", c.NewUserController().GetUserById)
	}

	return r
}
