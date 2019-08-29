package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
{
	"code": 0, //0表示成功,其他表示失败
	"message":"success"， //用来描述失败的原因
	"data":{

	}
}
*/

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//返回错误
func ResponseError(ctx *gin.Context, code int, message string) {

	responseData := &ResponseData{
		Code:    code,
		Message: message,
	}

	ctx.JSON(http.StatusOK, responseData) //渲染到客户端
}

//返回成功
func ResponseSuccess(ctx *gin.Context, message string, data interface{}) {
	responseData := &ResponseData{
		Code:    CodeSuccess,
		Message: message,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, responseData)
}
