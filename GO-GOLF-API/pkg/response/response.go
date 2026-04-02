package response

import (
	"GO-GOLF-API/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"result"`
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

func ErrorResponseWithMessage(c *gin.Context, statusCode int, appCode int, message string) {
	c.JSON(statusCode, ResponseData{
		Code:    appCode,
		Message: message,
		Data:    nil,
	})
}

type HandlerFunc func(ctx *gin.Context) (res interface{}, err error)

func Wrap(handler HandlerFunc) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		res, err := handler(ctx)
		if err != nil {
			HandleGlobalError(ctx, err)
			return
		}
		SuccessResponse(ctx, CodeSuccess, res)
	}
}

func HandleGlobalError(ctx *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		global.Logger.Warn("[APP_ERROR]",
			zap.Int("status", appErr.StatusCode),
			zap.Int("code", appErr.AppCode),
			zap.String("path", ctx.Request.URL.Path),
			zap.Error(appErr.RootErr),
		)

		if appErr.Message != "" {
			ErrorResponseWithMessage(ctx, appErr.StatusCode, appErr.AppCode, appErr.Message)
			return
		}

		ErrorResponse(ctx, appErr.StatusCode, appErr.AppCode)
	} else {
		global.Logger.Error("[SYSTEM_ERROR]",
			zap.String("path", ctx.Request.URL.Path),
			zap.Error(err),
		)

		ErrorResponse(ctx, http.StatusInternalServerError, ErrInternalError)
	}
}
