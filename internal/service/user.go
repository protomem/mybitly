package service

import (
	"github.com/protomem/mybitly/internal/dto"
	"github.com/protomem/mybitly/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	userRepo UserRepository
}

func NewUser(ur UserRepository) *User {
	return &User{
		userRepo: ur,
	}
}

func (u *User) GenerateHashPassword(password string) string {
	return ""
}

func (u *User) FindByNickname(nickname string) (model.User, error) {

	filter := bson.D{{Key: "nickname", Value: nickname}}
	return u.userRepo.Find(filter)

}

func (u *User) Create(user dto.UserCreate) (model.User, error) {

	if _, err := u.userRepo.Find(bson.D{{Key: "nickname", Value: user.Nickname}}); err != mongo.ErrNoDocuments {
		return model.User{}, err
	}

	if _, err := u.userRepo.Find(bson.D{{Key: "email", Value: user.Email}}); err != mongo.ErrNoDocuments {
		return model.User{}, err
	}

	userModel := model.User{
		Nickname:       user.Nickname,
		Password:       u.GenerateHashPassword(user.Password),
		Email:          user.Email,
		IsConfirmEmail: false,
	}

	if err := u.userRepo.Create(userModel); err != nil {
		return model.User{}, err
	}

	userFromDB, err := u.FindByNickname(user.Nickname)
	if err != nil {
		return model.User{}, err
	}

	return userFromDB, nil

}

func (u *User) DeleteByNickname(nickname string) error {

	filter := bson.D{{Key: "nickname", Value: nickname}}
	return u.userRepo.Delete(filter)

}
