package service

import "github.com/protomem/mybitly/internal/repository"

type Services struct {
	*LinkPair
}

func New(repos *repository.Repositories) *Services {
	return &Services{
		NewLinkPair(repos.LinkPair),
	}
}
