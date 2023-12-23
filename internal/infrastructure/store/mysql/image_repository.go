package mysql

import (
	"github.com/ningzining/cotton-pavilion/internal/domain/model"
	"github.com/ningzining/cotton-pavilion/internal/domain/repository"
	"gorm.io/gorm"
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
