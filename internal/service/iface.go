package service

import "github.com/protomem/mybitly/internal/model"

type (
	LinkPairRepository interface {
		FindAll(filter interface{}) ([]model.LinkPair, error)
		Find(filter interface{}) (model.LinkPair, error)
		Create(model.LinkPair) error
		Delete(filter interface{}) error
	}

	UserRepository interface {
		Find(filter interface{}) (model.User, error)
		Create(model.User) error
		Delete(filter interface{}) error
	}
)
