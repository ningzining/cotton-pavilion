package repository

import "user-center/internal/domain/model"

type ImageRepository interface {
	Save(image *model.Image) error
	FindById(id uint64) (*model.Image, error)
}
