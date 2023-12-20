package mysql

import (
	"gorm.io/gorm"
	"user-center/internal/domain/model"
	"user-center/internal/domain/repository"
)

type ImageRepository struct {
	db *gorm.DB
}

func NewImageRepository(r Repository) repository.ImageRepository {
	return &ImageRepository{db: r.DB}
}

func (r *ImageRepository) Save(image *model.Image) error {
	return r.db.Create(image).Error
}

func (r *ImageRepository) FindById(id uint64) (*model.Image, error) {
	var image model.Image
	err := r.db.Where("id = ?", id).Find(&image).Error
	return &image, err
}
