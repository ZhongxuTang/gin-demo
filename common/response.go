package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS = 200
	FAIL    = 500
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Result(code int, msg string, data interface{}, context *gin.Context) {
	context.JSONP(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func Success(context *gin.Context) {
	SuccessWithData(nil, context)
}

func SuccessWithData(data interface{}, context *gin.Context) {
	Result(SUCCESS, "success", data, context)
}

func SuccessWithMessage(message string, context *gin.Context) {
	Result(SUCCESS, message, nil, context)
}

func SuccessWithMessageData(message string, data interface{}, context *gin.Context) {
	Result(SUCCESS, message, data, context)
}

func Fail(context *gin.Context) {
	Result(FAIL, "fail", nil, context)
}

func FailWithMessage(message string, context *gin.Context) {
	Result(FAIL, message, nil, context)
}

func FailWithMessageData(message string, data interface{}, context *gin.Context) {
	Result(FAIL, message, data, context)
}
