package tags

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(tagsController *TagsController, router *gin.Engine) *gin.Engine {

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	// prefix -> /api
	baseRouter := router.Group("/api")
	tagsRouter := baseRouter.Group("/tags")

	tagsRouter.GET("", tagsController.FindAll)
	tagsRouter.GET("/:tagId", tagsController.FindById)
	tagsRouter.POST("", tagsController.Create)
	tagsRouter.PATCH("/:tagId", tagsController.Update)
	tagsRouter.DELETE("/:tagId", tagsController.Delete)

	return router
}
