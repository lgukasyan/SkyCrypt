package response

import (
	"github.com/gin-gonic/gin"
)

type ResponseDetails struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Method  string      `json:"method"`
	Url     string      `json:"url"`
	Status  int         `json:"status"`
}

func Write(ctx *gin.Context, statusCode int, data interface{}, msg string) {
	ctx.JSON(statusCode, &ResponseDetails{
		Data:    data,
		Message: msg,
		Method:  ctx.Request.Method,
		Url:     ctx.Request.URL.Path,
		Status:  statusCode,
	})
}
