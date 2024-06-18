package controller

import (
	"github.com/duynguyenbui/go-ecommerce-backend-api/internal/service"
	"github.com/duynguyenbui/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// controller -> service -> repo -> models -> dbs
func (uc *UserController) GetUserById(c *gin.Context) {
	response.SuccessResponse(c, 20001, []string{"duynguyenbui", "m10", "cr7"})
}
