package request

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetIdParam(ctx *gin.Context) uint64 {
	param := ctx.Param("id")
	id, _ := strconv.ParseUint(param, 10, 64)
	return id
}

func GetPageParam(ctx *gin.Context) PageParam {
	page, _ := strconv.Atoi(ctx.Query("page"))
	if page <= 0 {
		page = 1
	}

	size, _ := strconv.Atoi(ctx.Query("size"))
	if size <= 0 {
		size = 20
	}

	return PageParam{
		Page: page,
		Size: size,
	}
}
