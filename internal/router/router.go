package router

import (
	"github.com/gin-gonic/gin"
	"github.com/protomem/mybitly/internal/controller"
)

func New(controllers *controller.Controllers) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			controllers.LinkPair.Route("/linkPairs", v1)
		}
	}

	return router
}
