package application

import (
	"fmt"
	"path/filepath"
	"time"
	"user-center/internal/application/types"
	"user-center/internal/domain/model"
	"user-center/internal/infrastructure/store"
	"user-center/internal/infrastructure/third_party"
	"user-center/pkg/logger"
)

type ImageApplication interface {
	Upload(dto types.UploadDTO) (*types.UploadRet, error)
}

type imageApplication struct {
	Store store.Factory
	oss   third_party.Oss
}

func newImageApplication(store store.Factory, oss third_party.Oss) ImageApplication {
	return &imageApplication{
		Store: store,
		oss:   oss,
	}
}

func (i imageApplication) Upload(dto types.UploadDTO) (*types.UploadRet, error) {
	ext := filepath.Ext(dto.File.Filename)
	objectName := fmt.Sprintf("%s%d%s", "example/", time.Now().UnixMilli(), ext)
	file, err := dto.File.Open()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s", "https://cotton-pavilion.oss-cn-hangzhou.aliyuncs.com/", objectName)
	if err := i.oss.PutObject(objectName, file); err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	image := &model.Image{
		Url:          url,
		CreateUserId: dto.UserId,
	}
	if err := i.Store.ImageRepository().Save(image); err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return &types.UploadRet{
		Url: url,
	}, nil
}
