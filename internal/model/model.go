package model

import "github.com/protomem/mybitly/internal/types"

type Model struct {
	ID types.ID `bson:"_id"`

	CreatedAt types.Time `bson:"createdAt"`
	UpdatedAt types.Time `bson:"updatedAt"`
}
