package service

import (
	"time"

	"github.com/protomem/mybitly/internal/dto"
	"github.com/protomem/mybitly/internal/model"
	"github.com/protomem/mybitly/internal/types"
	"github.com/protomem/mybitly/pkg/crypt"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	_defautlTTL = 5 * time.Minute
)

type LinkPair struct {
	linkPairRepo LinkPairRepository
}

func NewLinkPair(lpr LinkPairRepository) *LinkPair {
	return &LinkPair{
		linkPairRepo: lpr,
	}
}

// FIXME Add Validity Check

func (lp *LinkPair) GenerateShortLink(originLink string) (string, error) {
	return crypt.GenerateShortToken(originLink)
}

func (lp *LinkPair) FindAll() ([]model.LinkPair, error) {
	return lp.linkPairRepo.FindAll(bson.D{{}})
}

func (lp *LinkPair) FindByShortLink(shortLink string) (model.LinkPair, error) {

	filter := bson.D{{Key: "shortLink", Value: shortLink}}
	return lp.linkPairRepo.Find(filter)

}

func (lp *LinkPair) Create(linkPair dto.LinkPairCreate) (model.LinkPair, error) {

	var err error

	shortLink, err := lp.GenerateShortLink(linkPair.OriginLink)
	if err != nil {
		return model.LinkPair{}, err
	}

	linkPairModel := model.LinkPair{
		ShortLink:  shortLink,
		OriginLink: linkPair.OriginLink,
		Exp:        types.Time(time.Now().Add(_defautlTTL).Unix()),
	}

	if err := lp.linkPairRepo.Create(linkPairModel); err != nil {
		return model.LinkPair{}, err
	}

	linkPairFromRepo, err := lp.FindByShortLink(shortLink)
	if err != nil {
		return model.LinkPair{}, err
	}

	return linkPairFromRepo, nil

}
