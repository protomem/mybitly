package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repositories struct {
	*LinkPair
}

func New(client *mongo.Client) *Repositories {
	db := client.Database("mybitly-db")

	return &Repositories{
		LinkPair: NewLinkPair(db),
	}
}
