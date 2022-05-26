package repository

import (
	"context"
	"time"

	"github.com/protomem/mybitly/internal/model"
	"github.com/protomem/mybitly/internal/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LinkPair struct {
	coll *mongo.Collection
}

func NewLinkPair(db *mongo.Database) *LinkPair {
	return &LinkPair{
		coll: db.Collection(model.LinkPairCollectionName),
	}
}

// TODO Implement all methhods

func (lp *LinkPair) FindAll(filter interface{}) ([]model.LinkPair, error) {

	var linkPairs []model.LinkPair

	cur, err := lp.coll.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(context.Background()) {
		var linkPair model.LinkPair
		if err := cur.Decode(&linkPair); err != nil {
			return linkPairs, err
		}

		linkPairs = append(linkPairs, linkPair)
	}

	if err := cur.Err(); err != nil {
		return linkPairs, err
	}

	cur.Close(context.Background())

	if len(linkPairs) == 0 {
		return nil, mongo.ErrNilDocument
	}

	return linkPairs, nil

}

func (lp *LinkPair) Find(filter interface{}) (model.LinkPair, error) {

	res := lp.coll.FindOne(context.Background(), filter)

	if err := res.Err(); err != nil {
		return model.LinkPair{}, err
	}

	var linkPair model.LinkPair

	if err := res.Decode(&linkPair); err != nil {
		return model.LinkPair{}, err
	}

	return linkPair, nil

}

func (lp *LinkPair) Create(linkPair model.LinkPair) error {

	currentTime := types.Time(time.Now().Unix())
	linkPair.CreatedAt = currentTime
	linkPair.UpdatedAt = currentTime

	linkPair.ID = types.ID(primitive.NewObjectID())

	_, err := lp.coll.InsertOne(context.Background(), linkPair)
	if err != nil {
		return err
	}

	return nil

}