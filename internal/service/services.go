package service

import "github.com/protomem/mybitly/internal/repository"

type Services struct {
	*LinkPair
	*User
}

func New(repos *repository.Repositories) *Services {
	return &Services{
		LinkPair: NewLinkPair(repos.LinkPair),
		User:     NewUser(repos.User),
	}
}
