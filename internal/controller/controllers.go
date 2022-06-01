package controller

import "github.com/protomem/mybitly/internal/service"

type Controllers struct {
	*LinkPair
	*User
}

func New(services *service.Services) *Controllers {
	return &Controllers{
		LinkPair: NewLinkPair(services.LinkPair),
		User:     NewUser(services.User),
	}
}
