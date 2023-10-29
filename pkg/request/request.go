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
	currentPageStr := ctx.Query("current_page")
	currentPage, _ := strconv.Atoi(currentPageStr)
	if currentPage <= 0 {
		currentPage = 1
	}

	pageSizeStr := ctx.Query("page_size")
	pageSize, _ := strconv.Atoi(pageSizeStr)
	if pageSize <= 0 {
		pageSize = 20
	}

	return PageParam{
		CurrentPage: currentPage,
		PageSize:    pageSize,
	}
}
