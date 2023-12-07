package response

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"user-center/pkg/code"
	"user-center/pkg/errors"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, Result{
		Code: code.ErrSuccess,
		Msg:  "",
		Data: data,
	})
}

func Error(ctx *gin.Context, err error) {
	coder := errors.ParseCoder(err)
	if coder != nil {
		log.Println(err)
		ctx.JSON(coder.HTTPStatus(), Result{
			Code: coder.Code(),
			Msg:  coder.String(),
			Data: nil,
		})
		return
	}
}

func AuthError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusUnauthorized, Result{
		Code: 401,
		Msg:  err.Error(),
		Data: nil,
	})
}
