package interfaces

import (
	"github.com/gin-gonic/gin"
	"user-center/internal/application"
	"user-center/internal/application/types"
	"user-center/internal/domain/service"
	"user-center/internal/infrastructure/consts"
	"user-center/internal/infrastructure/store"
	"user-center/internal/infrastructure/third_party"
	"user-center/internal/infrastructure/util/jwtutil"
	"user-center/pkg/response"
)

type ImageHandler struct {
	Application application.Application
}

func NewImageHandler(store store.Factory, service service.Service, oss third_party.Oss) *ImageHandler {
	return &ImageHandler{
		Application: application.NewApplication(store, service, oss),
	}
}

// Upload
//
//	@Summary	图片资源上传
//	@Tags		common
//	@Param		req	body		types.UploadDTO	true	"req"
//	@Success	200	{object}	response.Result
//	@Router		/v1/common/upload [post]
func (i ImageHandler) Upload(ctx *gin.Context) {
	var dto types.UploadDTO
	if err := ctx.Bind(&dto); err != nil {
		response.Error(ctx, err)
		return
	}
	user := ctx.MustGet(consts.ContextKeyUser).(jwtutil.User)
	dto.UserId = user.Id
	data, err := i.Application.ImageApplication().Upload(dto)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, data)
}
