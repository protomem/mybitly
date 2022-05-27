package model

import "github.com/protomem/mybitly/internal/types"

type Model struct {
	ID types.ID `bson:"_id" json:"-"`

	CreatedAt types.Time `bson:"createdAt"`
	UpdatedAt types.Time `bson:"updatedAt"`
}
