package repository

import (
	"context"
	"errors"
	"time"

	"github.com/protomem/mybitly/internal/model"
	"github.com/protomem/mybitly/internal/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	collections *mongo.Collection
}

func NewUser(db *mongo.Database) *User {
	return &User{
		collections: db.Collection(model.UserCollectionName),
	}
}

func (u *User) Find(filter interface{}) (model.User, error) {

	res := u.collections.FindOne(context.Background(), filter)

	if err := res.Err(); err != nil {
		return model.User{}, err
	}

	var user model.User

	if err := res.Decode(&user); err != nil {
		return model.User{}, err
	}

	return user, nil

}

func (u *User) Create(user model.User) error {

	currentTime := types.Time(time.Now().Unix())
	user.CreatedAt = currentTime
	user.UpdatedAt = currentTime

	user.ID = types.ID(primitive.NewObjectID())

	_, err := u.collections.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil

}

func (u *User) Delete(filter interface{}) error {

	res, err := u.collections.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("not user were deleted")
	}

	return nil

}
