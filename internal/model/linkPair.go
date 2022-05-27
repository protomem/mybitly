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

	ShortLink  string `bson:"shortLink" json:"shortLink"`
	OriginLink string `bson:"originLink" json:"originLink"`

	Exp types.Time `bson:"exp" json:"exp"`
}

func (lp *LinkPair) IsAlive() bool {

	currentTime := types.Time(time.Now().Unix())
	return lp.Exp > currentTime

}
