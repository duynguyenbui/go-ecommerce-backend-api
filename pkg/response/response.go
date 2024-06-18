package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`    // code status
	Message string      `json:"message"` // message
	Data    interface{} `json:"data"`    // data return
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusInternalServerError, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
