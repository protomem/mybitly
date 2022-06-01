package model

const (
	UserCollectionName = "users"
)

type User struct {
	Model

	Nickname string `bson:"nickname" json:"nickname"`
	Password string `bson:"password" json:"-"`

	Email          string `bson:"email"          json:"email"`
	IsConfirmEmail bool   `bson:"isConfirmEmail" json:"isConfirmEmail"`
}
