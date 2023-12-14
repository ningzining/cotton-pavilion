package store

import "user-center/internal/domain/repository"

var client Factory

type Factory interface {
	UserRepository() repository.UserRepository
	AutoMigrate()
	Clean()
}

func Client() Factory {
	return client
}

func SetClient(factory Factory) {
	client = factory
}
