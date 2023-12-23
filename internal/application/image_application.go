package application

import (
	"fmt"
	"github.com/ningzining/cotton-pavilion/internal/application/types"
	"github.com/ningzining/cotton-pavilion/internal/domain/model"
	"github.com/ningzining/cotton-pavilion/internal/infrastructure/store"
	"github.com/ningzining/cotton-pavilion/internal/infrastructure/third_party"
	"github.com/ningzining/cotton-pavilion/pkg/logger"
	"path/filepath"
	"time"
)

type ImageApplication interface {
	Upload(dto types.UploadDTO) (*types.UploadRet, error)
}

type imageApplication struct {
	Store store.Factory
}

func newImageApplication(store store.Factory) ImageApplication {
	return &imageApplication{
		Store: store,
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

	ossClient, err := third_party.NewOssClient(nil)
	if err != nil {
		return nil, err
	}
	if err := ossClient.PutObject(objectName, file); err != nil {
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
