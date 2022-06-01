package repository

import (
	"github.com/protomem/mybitly/pkg/mdb"
)

type Repositories struct {
	*LinkPair
	*User
}

func New(client *mdb.Client) *Repositories {
	db := client.Database("mybitly-db")

	return &Repositories{
		LinkPair: NewLinkPair(db),
		User:     NewUser(db),
	}
}
