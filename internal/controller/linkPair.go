package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/protomem/mybitly/internal/dto"
	"github.com/protomem/mybitly/internal/service"
)

const (
	_shortLinkParam = "shortLink"
)

type LinkPair struct {
	linkPairServ *service.LinkPair
}

func NewLinkPair(lps *service.LinkPair) *LinkPair {
	return &LinkPair{
		linkPairServ: lps,
	}
}

func (lp *LinkPair) Route(path string, rg *gin.RouterGroup) {

	linkPairs := rg.Group(path)
	{
		linkPairs.GET("/", lp.GetList)
		linkPairs.GET("/:shortLink", lp.Get)
		linkPairs.POST("/", lp.Create)
		linkPairs.DELETE("/:shortLink", lp.Delete)
		linkPairs.PUT("/:shortLink")
	}

}

func (lp *LinkPair) GetList(ctx *gin.Context) {

	linkPairs, err := lp.linkPairServ.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, linkPairs)

}

func (lp *LinkPair) Get(ctx *gin.Context) {

	shortLink := ctx.Param(_shortLinkParam)

	linkPair, err := lp.linkPairServ.FindByShortLink(shortLink)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, linkPair)

}

func (lp *LinkPair) Create(ctx *gin.Context) {

	var newLinkPair dto.LinkPairCreate

	if err := ctx.ShouldBindJSON(&newLinkPair); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	linkPair, err := lp.linkPairServ.Create(newLinkPair)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, linkPair)

}

func (lp *LinkPair) Delete(ctx *gin.Context) {

	shortLink := ctx.Param(_shortLinkParam)

	if err := lp.linkPairServ.DeleteByShortLink(shortLink); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)

}
