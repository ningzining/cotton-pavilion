package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, Result{
		Code: 0,
		Msg:  "",
		Data: data,
	})
}

func Error(ctx *gin.Context, err error) {
	errorWithMsg(ctx, err.Error())
}

func errorWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, Result{
		Code: 500,
		Msg:  msg,
		Data: nil,
	})
}
