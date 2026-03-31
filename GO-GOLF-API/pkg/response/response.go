package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, appCode int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    appCode,
		Message: msg[appCode],
		Data:    data,
	})
}

func SuccessResponseWithStatus(c *gin.Context, statusCode int, appCode int, data interface{}) {
	c.JSON(statusCode, ResponseData{
		Code:    appCode,
		Message: msg[appCode],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, appCode int) {
	c.JSON(statusCode, ResponseData{
		Code:    appCode,
		Message: msg[appCode],
		Data:    nil,
	})
}
