package store

import "github.com/ningzining/cotton-pavilion/internal/domain/repository"

var client Factory

type Factory interface {
	AutoMigrate()
	Clean()

	UserRepository() repository.UserRepository
	ImageRepository() repository.ImageRepository
}

func Client() Factory {
	return client
}

func SetClient(factory Factory) {
	client = factory
}
