package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/protomem/mybitly/internal/dto"
	"github.com/protomem/mybitly/internal/service"
)

const (
	_nicknameParam = "nickname"
)

type User struct {
	userServ *service.User
}

func NewUser(us *service.User) *User {
	return &User{
		userServ: us,
	}
}

func (u *User) Route(path string, rg *gin.RouterGroup) {

	users := rg.Group(path)
	{
		users.GET("/")
		users.GET("/:nickname", u.Get)
		users.POST("/", u.Create)
		users.DELETE("/:nickname", u.Delete)

		users.PATCH("/:nickname/email")
		users.PATCH("/:nickname/password")
	}

}

func (u *User) Get(ctx *gin.Context) {

	nickname := ctx.Param(_nicknameParam)

	user, err := u.userServ.FindByNickname(nickname)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)

}

func (u *User) Create(ctx *gin.Context) {

	var newUser dto.UserCreate

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := u.userServ.Create(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)

}

func (u *User) Delete(ctx *gin.Context) {

	nickname := ctx.Param(_nicknameParam)

	if err := u.userServ.DeleteByNickname(nickname); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)

}
