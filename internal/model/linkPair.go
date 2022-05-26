package model

import (
	"time"

	"github.com/protomem/mybitly/internal/types"
)

const (
	LinkPairCollectionName = "linkPairs"
)

type LinkPair struct {
	Model

	ShortLink  string `bson:"shortLink"`
	OriginLink string `bson:"originLink"`

	Exp types.Time `bson:"exp"`
}

func (lp *LinkPair) IsValid() bool {

	currentTime := types.Time(time.Now().Unix())
	return lp.Exp < currentTime

}
