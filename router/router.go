package router

import (
	"github.com/gin-gonic/gin"
	"hacktiv8_final/controller"
	"hacktiv8_final/posts"
	"net/http"
)

func NewRouter(commentRepo posts.RepositoryComment, commentController *controller.PostsController) *gin.Engine {
	router := gin.Default()
	router.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	groupRouter := router.Group("/posts")
	groupRouter.POST("/comment", commentController.CreateComment)
	groupRouter.GET("/comments", commentController.FindAllComment)
	groupRouter.GET("/comment/:id", commentController.FindByIdComment)
	groupRouter.DELETE("/comment/:id", commentController.DeleteComment)
	groupRouter.PUT("/comment", commentController.UpdateComment)
	return router
}
