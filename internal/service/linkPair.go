package service

import (
	"errors"
	"time"

	"github.com/protomem/mybitly/internal/dto"
	"github.com/protomem/mybitly/internal/model"
	"github.com/protomem/mybitly/internal/types"
	"github.com/protomem/mybitly/pkg/crypt"
	u "github.com/rjNemo/underscore"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	_defautlTTL = 5 * 24 * time.Hour
)

type LinkPair struct {
	linkPairRepo LinkPairRepository
}

func NewLinkPair(lpr LinkPairRepository) *LinkPair {
	return &LinkPair{
		linkPairRepo: lpr,
	}
}

func (lp *LinkPair) GenerateShortLink(originLink string) (string, error) {
	return crypt.GenerateShortToken(originLink)
}

func (lp *LinkPair) FindAll() ([]model.LinkPair, error) {

	filter := bson.D{{}}
	linkPairs, err := lp.linkPairRepo.FindAll(filter)
	if err != nil {
		return nil, err
	}

	aliveLinkPairs := u.Filter(linkPairs, func(n model.LinkPair) bool { return n.IsAlive() })
	if len(aliveLinkPairs) == 0 {
		return nil, errors.New("no live short links")
	}

	return aliveLinkPairs, nil

}

func (lp *LinkPair) FindByShortLink(shortLink string) (model.LinkPair, error) {

	filter := bson.D{{Key: "shortLink", Value: shortLink}}
	linkPair, err := lp.linkPairRepo.Find(filter)
	if err != nil {
		return linkPair, err
	}

	if !linkPair.IsAlive() {
		return linkPair, errors.New("this short link is dead")
	}

	return linkPair, nil

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

func (lp *LinkPair) DeleteByShortLink(shortLink string) error {

	filter := bson.D{{Key: "shortLink", Value: shortLink}}
	return lp.linkPairRepo.Delete(filter)

}
