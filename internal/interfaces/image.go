package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/ningzining/cotton-pavilion/internal/application"
	"github.com/ningzining/cotton-pavilion/internal/application/types"
	"github.com/ningzining/cotton-pavilion/internal/domain/service"
	"github.com/ningzining/cotton-pavilion/internal/infrastructure/consts"
	"github.com/ningzining/cotton-pavilion/internal/infrastructure/store"
	"github.com/ningzining/cotton-pavilion/internal/infrastructure/util/jwtutil"
	"github.com/ningzining/cotton-pavilion/pkg/response"
)

type ImageHandler struct {
	Application application.Application
}

func NewImageHandler(store store.Factory, service service.Service) *ImageHandler {
	return &ImageHandler{
		Application: application.NewApplication(store, service),
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
